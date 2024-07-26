import SearchModal from '@/components/searchModal';
import SongList from '@/components/SongList';
import { useState } from 'react';
import { Button } from './components/ui/button';
import PlayList from './components/PlayList';

function App() {
  const [path, setPath] = useState<'playlist' | 'list'>('list');

  return (
    <div className="w-full flex flex-col gap-10 items-center justify-start py-10 px-6 flex-1">
      <div className="w-full flex justify-between px-6">
        <h1 className="text-2xl font-bold uppercase text-start">
          Search for name of the song
        </h1>
        <div className="flex gap-4">
          <Button variant={path === 'playlist' ? 'default' : 'outline'} onClick={() => setPath('playlist')}>
            PlayList
          </Button>
          <Button variant={path === 'list' ? 'default' : 'outline'} onClick={() => setPath('list')}>
            List
          </Button>
          <SearchModal />
        </div>
      </div>
      {
        path === 'playlist' ? <PlayList /> : <SongList />
      }
    </div>
  );
}

export default App;
