import { apiSlice } from '@/api/api-slice';

export const customerSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    createNewCustomer: builder.mutation({
      query: ({
        full_name,
        email,
        phone_number,
        marital_status,
        employment_status,
        name_of_employer,
        date_of_birth,
        identification_card,
        identification_card_type,
        address,
      }) => ({
        url: '/customer/create',
        method: 'POST',
        body: {
          full_name,
          email,
          phone_number,
          marital_status,
          employment_status,
          name_of_employer,
          date_of_birth,
          identification_card,
          identification_card_type,
          address,
          phone_number,
        },
      }),
    }),
    getAllCustomers: builder.mutation({
      query: ({ email, password }) => ({
        url: '/customer',
        method: 'GET',
        body: {
          email: email,
          password: password,
        },
      }),
    }),
  }),
});

export const { useCreateNewCustomerMutation, use } = customerSlice;
