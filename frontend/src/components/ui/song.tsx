import { SongType } from '@/types/song';
import { convertirMilisegundos } from '@/lib/utils';
import { PlusIcon } from 'lucide-react';
import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
  AlertDialogAction,
} from '@/components/ui/alert-dialog';
import { Button } from './button';
import { usePlayList } from '../providers/PlayListProvider';

export default function Song({ song, index }: { song: SongType, index: number }) {
  const { putSong } = usePlayList();

  return (
    <div className="flex border-t border-border py-3 px-6 gap-4 items-center w-full">
      <AlertDialog>
        <AlertDialogTrigger asChild>
          <button type="button" className="flex py-3 px-6 gap-4 items-center w-full">
            <h1 className="overflow-x-hidden text-nowrap w-full text-start">{song.TrackName}</h1>
            <h2 className="w-full max-w-64 overflow-x-hidden text-nowrap text-start">{song.ArtistName}</h2>
            <h2 className="w-full max-w-24 text-center">{song.Popularity}</h2>
            <h2 className="w-full max-w-24 text-end">{song.Year}</h2>
            <h2 className="w-full max-w-24 flex justify-end items-center">{convertirMilisegundos(song.DurationMs)}</h2>
          </button>
        </AlertDialogTrigger>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>{song.TrackName}</AlertDialogTitle>
            <AlertDialogDescription className="flex flex-col gap-2">
              <span>
                Artista:
                {' '}
                {song.ArtistName}
              </span>
              <span>
                Duracion en milisegundos:
                {' '}
                {song.DurationMs}
              </span>
              <span>
                Popularidad:
                {' '}
                {song.Popularity}
              </span>
              <span>
                Año:
                {' '}
                {song.Year}
              </span>
              <span>
                TrackId:
                {' '}
                {song.TrackId}
              </span>
              <span>
                Genero:
                {' '}
                {song.Genre}
              </span>
              <span>
                Valance:
                {' '}
                {song.Valence}
              </span>
            </AlertDialogDescription>
            <AlertDialogFooter>
              <AlertDialogAction>Ok</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogHeader>
        </AlertDialogContent>
      </AlertDialog>
      <div className="w-full max-w-16 flex justify-end items-center">
        <Button
          variant="outline"
          size="icon"
          onClick={() => {
            putSong(index);
          }}
        >
          <PlusIcon className="w-4" />
        </Button>
      </div>
    </div>

  );
}
