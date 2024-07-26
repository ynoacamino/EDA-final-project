import { useEffect, useState } from 'react';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { usePlayList } from './providers/PlayListProvider';
import PlayListSong from './ui/PlayListSong';
import { Button } from './ui/button';

export default function PlayList() {
  const {
    playlist, randomPlayList, clearPlayList, orderPlayList,
  } = usePlayList();

  const [order, setOrder] = useState(1);
  const [filter, setFilter] = useState(1);

  useEffect(() => {
    orderPlayList(filter, order);
  }, [order, filter]);

  return (
    <div className="w-full max-w-6xl border border-border rounded-xl min-h-[70vh]">
      <div className="w-full items-center justify-between flex">
        <h1 className="text-2xl font-bold text-start my-4 px-4">
          PlayList
        </h1>
        <div className="flex items-center gap-4">
          <Select onValueChange={(e) => setFilter(Number(e))}>
            <SelectTrigger className="w-[150px]">
              <SelectValue placeholder="Select a filter" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel>Filter</SelectLabel>
                <SelectItem value="1">Year</SelectItem>
                <SelectItem value="2">Duration</SelectItem>
                <SelectItem value="3">Popurarity</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
          <Select onValueChange={(e) => setOrder(Number(e))}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Select a order" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel>Order</SelectLabel>
                <SelectItem value="1">Asc</SelectItem>
                <SelectItem value="-1">Des</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
          <Button
            variant="outline"
            onClick={() => {
              randomPlayList();
            }}
          >
            Shuffle
          </Button>
          <Button variant="outline" onClick={() => clearPlayList()}>
            Clear
          </Button>
        </div>
      </div>
      <div className="flex border-t border-border py-3 px-6 gap-4 text-muted-foreground">
        <span className="overflow-x-hidden text-nowrap w-full max-w-8" />
        <span className="overflow-x-hidden text-nowrap w-full">
          Track Name
        </span>
        <span className="w-full max-w-64 overflow-x-hidden text-nowrap">
          Artist Name
        </span>
        <span className="w-full max-w-24 text-center">
          Popularity
        </span>
        <span className="w-full max-w-24 text-end">
          Year
        </span>
        <span className="w-full max-w-24 flex justify-end items-center">
          Duration
        </span>
        <span className="w-full max-w-16 flex justify-end items-center">
          Action
        </span>
      </div>
      {
      playlist.map((song, index) => (
        <PlayListSong key={crypto.randomUUID()} song={song} index={index} />
      ))
    }
    </div>
  );
}
