import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const convertirMilisegundos = (minseg: number):string => {
  // Convierte la cantidad de milisegundos a segundos
  let segundos = Math.floor(minseg / 1000);

  // Calcula los minutos y los segundos restantes
  const minutos = Math.floor(segundos / 60);
  segundos %= 60;

  // Asegúrate de que los segundos siempre tengan dos dígitos
  const segundosStr = segundos < 10 ? `0${segundos}` : segundos;

  return `${minutos}:${segundosStr}`;
};
