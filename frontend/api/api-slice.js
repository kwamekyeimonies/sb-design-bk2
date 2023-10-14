import { fetchBaseQuery, createApi } from '@reduxjs/toolkit/query/react';
import Cookies from 'js-cookie';

const baseURL = 'http://192.168.8.145:9098/api/v1';

const baseQuery = fetchBaseQuery({
  baseUrl: baseURL,
  // prepareHeaders: (headers) => {
  //     const token = Cookies.get('token');
  //     if (token) {
  //         headers.set('Authorization', `Bearer ${token}`)
  //     }
  //     return headers
  // }
});

export const apiSlice = createApi({
  reducerPath: 'api',
  baseQuery: baseQuery,
  endpoints: (builder) => ({}),
});
