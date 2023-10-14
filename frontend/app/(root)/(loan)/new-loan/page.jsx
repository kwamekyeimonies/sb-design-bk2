import { NewLoanForm } from '@/components';
import { Box, Typography } from '@mui/material';
import React from 'react';

const NewLoan = () => {
  return (
    <Box className='w-2/4'>
      <Typography variant='h4' className='text-center font-bold mb-5 uppercase'>
        New Loan
      </Typography>
      <NewLoanForm />
    </Box>
  );
};

export default NewLoan;
