'use client';
import authenticationReducer from '@/api/authenticationReducer';
import { customerSlice } from '@/slices/customer-slice';
import { loanSlice } from '@/slices/loan-slice';
import { authenticationSlice } from '@/slices/auth-slice';
import { configureStore } from '@reduxjs/toolkit';

const store = configureStore({
  reducer: {
    auth: authenticationReducer,
    [authenticationSlice.reducerPath]: authenticationSlice.reducer,
    [customerSlice.reducerPath]: customerSlice.reducer,
    [loanSlice.reducerPath]: loanSlice.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(authenticationSlice.middleware),
});

export default store;
