'use client';

import React, { useEffect, useState } from 'react';
import {
  TextField,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Button,
  Grid,
  Alert,
} from '@mui/material';
import { useCreateNewCustomerMutation } from '@/slices/customer-slice';
import { useRouter } from 'next/navigation';

const MaritalStatusOptions = ['Single', 'Married', 'Divorced', 'Widowed'];
const EmploymentStatusOptions = [
  'Employed',
  'Unemployed',
  'Self-employed',
  'Student',
];

const NewCustomerForm = () => {
  const router = useRouter();
  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    maritalStatus: '',
    employmentStatus: '',
    employerName: '',
    dob: '',
    email: '',
    idCardNumber: '',
    idCardType: '',
    address: '',
    phoneNumber: '',
  });

  const [createNewCustomer, { isLoading, isSuccess, isError, error, data }] =
    useCreateNewCustomerMutation();

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
    createNewCustomer({
      full_name: `${formData.firstName + ' ' + formData.lastName}`,
      phone_number: formData.phoneNumber,
      marital_status: formData.maritalStatus,
      employment_status: formData.employmentStatus,
      name_of_employer: formData.employerName,
      date_of_birth: formData.dob,
      identification_card: formData.idCardNumber,
      identification_card_type: formData.identification_card_type,
      address: formData.address,
      email: formData.email,
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
        <Grid item xs={6}>
          <TextField
            fullWidth
            label='First Name'
            name='firstName'
            value={formData.firstName}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={6}>
          <TextField
            fullWidth
            label='Last Name'
            name='lastName'
            value={formData.lastName}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={6}>
          <FormControl fullWidth>
            <InputLabel>Marital Status</InputLabel>
            <Select
              label='Marital Status'
              name='maritalStatus'
              value={formData.maritalStatus}
              onChange={handleChange}
              required
            >
              {MaritalStatusOptions.map((status) => (
                <MenuItem key={status} value={status}>
                  {status}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={6}>
          <FormControl fullWidth>
            <InputLabel>Employment Status</InputLabel>
            <Select
              label='Employment Status'
              name='employmentStatus'
              value={formData.employmentStatus}
              onChange={handleChange}
              required
            >
              {EmploymentStatusOptions.map((status) => (
                <MenuItem key={status} value={status}>
                  {status}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <FormControl fullWidth>
            <InputLabel>Email</InputLabel>
            <TextField
              fullWidth
              label='Email'
              name='email'
              value={formData.email}
              onChange={handleChange}
            />
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label='Name of Employer'
            name='employerName'
            value={formData.employerName}
            onChange={handleChange}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            type='date'
            label=''
            name='dob'
            value={formData.dob}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={6}>
          <TextField
            fullWidth
            label='ID Card Type'
            name='idCardType'
            value={formData.idCardType}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={6}>
          <TextField
            fullWidth
            label='ID Card Number'
            name='idCardNumber'
            value={formData.idCardNumber}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label='Address'
            name='address'
            value={formData.address}
            onChange={handleChange}
            required
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label='Phone Number. Eg. 0541234567'
            name='phoneNumber'
            value={formData.phoneNumber}
            onChange={handleChange}
            required
            // inputProps={{
            //   pattern: {
            //     value: /^(055|024|054|050|020|057|027)\d{7}$/,
            //     message: 'Invalid phone number format',
            //   },
            // }}
          />
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

export default NewCustomerForm;
