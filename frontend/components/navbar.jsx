import { Button } from '@mui/base';
import {
  AppBar,
  Box,
  IconButton,
  Menu,
  Toolbar,
  Typography,
} from '@mui/material';
import Link from 'next/link';
import React from 'react';

const Navbar = () => {
  return (
    <AppBar position='static'>
      <Toolbar className='flex items-center justify-between'>
        <Box className='flex items-center'>
          <Typography variant='h6' component='div'>
            SBDK
          </Typography>

          <Box className='flex'>
            <Typography className='mx-4'>
              <Link href='/dashboard'>Dashboard</Link>
            </Typography>
            <Typography className='mx-4'>
              <Link href='/new-customer'>New Customer</Link>
            </Typography>
            <Typography className='mx-4'>
              <Link href='/new-loan'>New Loan</Link>
            </Typography>
          </Box>
        </Box>

        <Button color='inherit'>Logout</Button>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
