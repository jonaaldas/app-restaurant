import {ofetch} from 'ofetch';
import { Restaurant } from '@/types/restaurants';

const api = ofetch.create({
  baseURL: 'http://localhost:3000/api',
})

interface SearchParams {
  query: string;
}

interface SearchResponse {
  data: Restaurant[];
}

export const searchRestaurants = async (params: SearchParams): Promise<Restaurant[]> => {
  const response: SearchResponse = await api('/search', {
    query: {
      query: params.query,
    },
  })
  return response.data;
}