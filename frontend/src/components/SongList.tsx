import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { useEffect, useState } from 'react';

import { ResultType } from '@/types/result';
import Song from '@/components/ui/song';
import {
  OrderByYear,
  OrderByPopularity,
  OrderByDuration,
} from '../../wailsjs/go/main/App';

export default function SongList() {
  const initResult: ResultType = {
    Songs: [],
    TimeLapse: 0,
    Size: 0,
  };

  const [resultByOrder, setResultByOrder] = useState<ResultType>(initResult);

  const [order, setOrder] = useState(1);
  const [page, setPage] = useState(0);
  const [filter, setFilter] = useState(1);

  useEffect(() => {
    if (filter === 1) {
      OrderByYear(order, page).then((r) => {
        setResultByOrder(r as ResultType);
      });
    } else if (filter === 2) {
      OrderByDuration(order, page).then((r) => {
        setResultByOrder(r as ResultType);
      });
    } else if (filter === 3) {
      OrderByPopularity(order, page).then((r) => {
        setResultByOrder(r as ResultType);
      });
    }
  }, [order, filter, page]);
  return (
    <div className="w-full max-w-6xl border border-border rounded-xl min-h-[70vh]">
      <div className="w-full items-center justify-between flex">
        <h1 className="text-2xl font-bold text-start my-4 px-4">
          Songs
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
          <Select onValueChange={(e) => setPage(Number(e))}>
            <SelectTrigger className="w-[150px]">
              <SelectValue placeholder="Select a page" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel>Page</SelectLabel>
                {
                  Array.from({ length: 100 }, (_, i) => (
                    <SelectItem key={i} value={`${i}`}>{i + 1}</SelectItem>
                  ))
                }
                {
                  Array.from({ length: 100 }, (_, i) => {
                    const num = Math.floor(resultByOrder.Size / 200) - 100;
                    return <SelectItem key={num + i} value={`${num + i}`}>{num + i + 1}</SelectItem>;
                  })
                }
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
        </div>
      </div>
      <div className="flex border-t border-border py-3 px-6 gap-4 text-muted-foreground">
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
        resultByOrder.Songs.map((song, index) => (
          <Song key={song.TrackId} song={song} index={index} />
        ))
      }
    </div>
  );
}
