import { SongType } from '@/types/song';
import { cn } from '@/lib/utils';
import Song from '../ui/song';

export default function Browser({
  result, title, timeLapse, size,
}: { result: SongType[], title: string, timeLapse: number, size: number }) {
  return (
    <div className={cn('w-full', {
      'border-l border-border': title === 'Trie',
    })}
    >
      <h1 className="text-2xl font-bold w-full text-center my-4">{title}</h1>
      <h2 className="text-lg  w-full text-center my-4 text-muted-foreground">
        {size}
        {' '}
        canciones encontradas en
        {' '}
        {timeLapse}
        {' '}
        ms
      </h2>
      <div className="grid grid-cols-6 py-3 px-6 text-muted-foreground">
        <span className="col-span-4">
          Track Name
        </span>
        <span>
          Artist Name
        </span>
        <span className="flex justify-end items-center">
          Popularity
        </span>
      </div>
      {result.map((song, index) => (
        <Song key={song.TrackId} song={song} index={index} />
      ))}
    </div>
  );
}
