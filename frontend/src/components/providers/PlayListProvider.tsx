/* eslint-disable react/jsx-no-constructed-context-values */

'use client';

import {
  useContext, createContext,
  useState,
  useEffect,
} from 'react';

import { SongType } from '@/types/song';
import {
  PutSong,
  GetPlayList,
  RandomPlayList,
  RemoveSong,
  ClearPlayList,
  OrderPlayList,
} from '../../../wailsjs/go/main/App';

const userContext = createContext({
  playlist: [] as SongType[],
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  putSong: (index: number) => {},
  randomPlayList: () => {},
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  removeSong: (index: number) => {},
  clearPlayList: () => {},
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  orderPlayList: (filter: number, order: number) => {},
});

export const usePlayList = () => {
  const contex = useContext(userContext);
  if (!useContext) throw new Error('contect not found');
  return contex;
};

export function PlayListProvider({ children }: { children: React.ReactNode }) {
  const [playlist, setPlayList] = useState<SongType[]>([]);

  const getPlayList = async () => {
    GetPlayList()
      .then((r) => {
        console.log(r.length);
        setPlayList(r as SongType[]);
      });
  };

  const putSong = async (index: number) => {
    PutSong(index)
      .then(() => getPlayList());
  };

  const randomPlayList = async () => {
    RandomPlayList()
      .then(() => getPlayList());
  };

  const removeSong = async (index: number) => {
    RemoveSong(index)
      .then(() => getPlayList());
  };

  const clearPlayList = async () => {
    ClearPlayList()
      .then(() => getPlayList());
  };

  const orderPlayList = async (filter: number, order: number) => {
    OrderPlayList(filter, order)
      .then(() => getPlayList());
  };

  useEffect(() => {
    getPlayList();
  }, []);

  return (
    <userContext.Provider
      value={{
        playlist,
        putSong,
        randomPlayList,
        removeSong,
        clearPlayList,
        orderPlayList,
      }}
    >
      {children}
    </userContext.Provider>
  );
}
