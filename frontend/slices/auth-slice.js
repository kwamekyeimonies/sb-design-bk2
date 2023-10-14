import { apiSlice } from '@/api/api-slice';

export const authenticationSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    createUserAccount: builder.mutation({
      query: ({
        full_name,
        password,
        phone_number,
        dob,
        bank_branch,
        email,
      }) => ({
        url: '/users/signup',
        method: 'POST',
        body: {
          email,
          full_name,
          phone_number,
          dob,
          bank_branch,
          password,
        },
      }),
    }),
    loginUser: builder.mutation({
      query: ({ email, password }) => ({
        url: '/users/login',
        method: 'POST',
        body: {
          email: email,
          password: password,
        },
      }),
    }),
  }),
});

export const { useCreateUserAccountMutation, useLoginUserMutation } =
  authenticationSlice;
