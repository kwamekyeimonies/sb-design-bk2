import React from 'react';
import { Box, Typography } from '@mui/material';
import { NewCustomerForm } from '@/components';

const NewCustomer = () => {
  return (
    <Box className='w-2/4'>
      <Typography variant='h4' className='text-center font-bold mb-5 uppercase'>
        New Customer
      </Typography>
      <NewCustomerForm />
    </Box>
  );
};

export default NewCustomer;
