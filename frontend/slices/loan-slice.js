import { apiSlice } from '@/api/api-slice';

export const loanSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    createNewLoan: builder.mutation({
      query: ({ principal, rate, loan_time, interest, user_id, id }) => ({
        url: '/customer/loan/create',
        method: 'POST',
        body: {
          principal,
          rate,
          loan_time,
          interest,
          user_id,
          id,
        },
      }),
    }),
    getAllLoans: builder.query({
      query: () => ({
        url: '/customer/loans',
        method: 'GET',
      }),
    }),
  }),
});

export const { useCreateNewLoanMutation, useGetAllLoansQuery } = loanSlice;
