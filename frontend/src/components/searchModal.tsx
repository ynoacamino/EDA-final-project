/* eslint-disable import/no-extraneous-dependencies */
import {
  MagnifyingGlassIcon,
} from '@radix-ui/react-icons';

import { useState, useEffect } from 'react';

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { ResultType } from '@/types/result';
import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';

import { convertirMilisegundos } from '@/lib/utils';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {
  SearchSongInIndexInvert,
  SearchSongInTrie,
} from '../../wailsjs/go/main/App';
import { usePlayList } from './providers/PlayListProvider';

const initResult: ResultType = {
  Songs: [],
  TimeLapse: 0,
  Size: 0,
};

export default function SearchModal() {
  const [open, setOpen] = useState(false);
  const [input, setInput] = useState('');

  const { putSong } = usePlayList();
  const [browser, setBrowser] = useState(1);

  const [resultInvertIndex, setResultInvertIndex] = useState<ResultType>(initResult);

  useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === 'j' && (e.metaKey || e.ctrlKey)) {
        e.preventDefault();
        setOpen((op) => !op);
      }
    };

    document.addEventListener('keydown', down);
    return () => document.removeEventListener('keydown', down);
  }, []);

  const handleSearch = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (input.trim() === '') {
      setResultInvertIndex(initResult);
      return;
    }

    let result;
    if (browser === 1)result = await SearchSongInTrie(input) as ResultType;
    else result = await SearchSongInIndexInvert(input) as ResultType;

    setResultInvertIndex(result);
  };

  return (
    <>
      <Button variant="outline" size="sm" className="justify-start" onClick={() => setOpen(true)}>
        <MagnifyingGlassIcon className="w-5 h-5 md:mr-4" />
        <span className="min-w-32 text-start hidden md:flex">
          Buscar productos...
        </span>
      </Button>
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="max-w-5xl bg-accent">
          <DialogHeader>
            <div className="w-full flex justify-between items-center">
              <DialogTitle>Busca tu producto favorito</DialogTitle>
              <div className="flex gap-8 items-center">
                <Select onValueChange={(e) => setBrowser(Number(e))}>
                  <SelectTrigger className="w-[150px]">
                    <SelectValue placeholder="Tree" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectGroup>
                      <SelectLabel>Select</SelectLabel>
                      <SelectItem value="1">Tree</SelectItem>
                      <SelectItem value="2">Invert indexes</SelectItem>
                    </SelectGroup>
                  </SelectContent>
                </Select>
                <span>
                  Time:
                  {' '}
                  {resultInvertIndex.TimeLapse}
                  ms
                </span>
                <span>
                  Size:
                  {' '}
                  {resultInvertIndex.Size}
                </span>
              </div>
            </div>
          </DialogHeader>
          <div className="flex flex-col gap-4">
            <form onSubmit={handleSearch} className="border-primary rounded-sm border bg-background flex h-14 items-center">
              <MagnifyingGlassIcon className="w-10 h-10 p-1" />
              <input type="text" name="" id="" className="appearance-none bg-transparent border-0 flex-1 text-xl h-full focus-visible:outline-none pl-1 text-primary" value={input} onChange={(e) => setInput(e.target.value)} />
            </form>
            <div className="flex flex-col gap-2">
              <div className="grid grid-cols-7 py-3 px-6 text-muted-foreground gap-4">
                <span className="col-span-3">
                  Name
                </span>
                <span>
                  Artist
                </span>
                <span className="w-full text-end">
                  Year
                </span>
                <span className="flex justify-end items-center">
                  Duration
                </span>
                <span className="flex justify-end items-center">
                  Action
                </span>
              </div>
              <ScrollArea className="max-h-96">
                {
                  resultInvertIndex.Songs.map((song, index) => (
                    <div className="grid grid-cols-7 border-t border-border py-3 px-6 gap-4" key={song.TrackId}>
                      <h1 className="col-span-3 overflow-x-hidden text-nowrap">{song.TrackName}</h1>
                      <h2 className="overflow-x-hidden text-nowrap">{song.ArtistName}</h2>
                      <h2 className="w-full text-end">{song.Year}</h2>
                      <h2 className="w-full flex justify-end items-center">{convertirMilisegundos(song.DurationMs)}</h2>
                      <div className="flex justify-end items-center">
                        <Button
                          variant="outline"
                          size="icon"
                          onClick={() => {
                            putSong(index);
                          }}
                        >
                          +
                        </Button>
                      </div>
                    </div>
                  ))
                }
              </ScrollArea>
            </div>
          </div>
        </DialogContent>
      </Dialog>
    </>

  );
}
