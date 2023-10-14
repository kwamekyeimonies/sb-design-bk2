'use client';

import React, { useEffect, useState } from 'react';
import {
  TextField,
  FormControl,
  InputLabel,
  Button,
  Grid,
  Alert,
} from '@mui/material';
import { useCreateNewLoanMutation } from '@/slices/loan-slice';
import { useRouter } from 'next/navigation';
import Cookies from 'js-cookie';

const NewLoanForm = () => {
  const router = useRouter();
  const [formData, setFormData] = useState({
    principal: '',
    rate: '',
    loanTime: '',
    interest: '',
    accountNumber: '',
  });

  const [createNewLoan, { isLoading, isSuccess, isError, error, data }] =
    useCreateNewLoanMutation();

  useEffect(() => {
    if (isSuccess) {
      router.push('/dashboard');
    }
  }, [isSuccess]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    createNewLoan({
      principal: formData.principal,
      rate: formData.rate,
      loan_time: formData.loanTime,
      interest: formData.interest,
      id: formData.accountNumber,
      user_id: Cookies.get('id'),
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      {isSuccess && (
        <Alert severity='success' sx={{ marginBottom: 2 }}>
          {data.message}
        </Alert>
      )}
      {isError && (
        <Alert severity='error' sx={{ marginBottom: 2 }}>
          {error?.data?.error}
        </Alert>
      )}
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label='Customer  Number'
            name='accountNumber'
            value={formData.accountNumber}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label='Principal'
            name='principal'
            value={formData.principal}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label='Rate'
            name='rate'
            value={formData.rate}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={12}>
          <InputLabel>Loan Time (Years)</InputLabel>
          <TextField
            fullWidth
            label='Eg. 5'
            name='loanTime'
            value={formData.loanTime}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={12}>
          <FormControl fullWidth>
            <TextField
              fullWidth
              label='Interest'
              name='interest'
              value={formData.interest}
              onChange={handleChange}
            />
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <Button
            variant='contained'
            color='primary'
            style={{ backgroundColor: '#1976d2' }}
            type='submit'
            className='w-full'
          >
            Submit
          </Button>
        </Grid>
      </Grid>
    </form>
  );
};

export default NewLoanForm;
