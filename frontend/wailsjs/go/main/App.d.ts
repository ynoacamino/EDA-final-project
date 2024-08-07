// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from '../models';

export function ClearPlayList():Promise<void>;

export function GetPlayList():Promise<Array<types.Song>>;

export function Invert(arg1:Array<types.Song>):Promise<Array<types.Song>>;

export function OrderByDuration(arg1:number,arg2:number):Promise<types.Result>;

export function OrderByPopularity(arg1:number,arg2:number):Promise<types.Result>;

export function OrderByYear(arg1:number,arg2:number):Promise<types.Result>;

export function OrderPlayList(arg1:number,arg2:number):Promise<void>;

export function PutSong(arg1:number):Promise<void>;

export function RandomPlayList():Promise<Array<types.Song>>;

export function RemoveSong(arg1:number):Promise<void>;

export function SearchSongInIndexInvert(arg1:string):Promise<types.Result>;

export function SearchSongInTrie(arg1:string):Promise<types.Result>;
