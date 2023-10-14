import { LoansList } from '@/components';
import { Box, Typography } from '@mui/material';
import React from 'react';

const Dashboard = () => {
  return (
    <Box>
      <Typography variant='h4' className='text-center font-bold mb-5 uppercase'>
        Loan List
      </Typography>
      <LoansList />
    </Box>
  );
};

export default Dashboard;
