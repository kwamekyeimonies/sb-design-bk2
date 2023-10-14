'use client';
import React, { useState, useEffect } from 'react';
import { Alert } from '@mui/material';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import CircularProgress from '@mui/material/CircularProgress';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { redirect } from 'next/navigation';
import { useLoginUserMutation } from '@/slices/auth-slice';
import { useDispatch } from 'react-redux';
import { authenticate } from '@/api/authenticationReducer';

const defaultTheme = createTheme();

export default function SignIn() {
  const dispatch = useDispatch();

  const [loginUser, { data, isLoading, isError, error, isSuccess }] =
    useLoginUserMutation();
  const [errors, setErrors] = useState({
    email: '',
    password: '',
  });
  const [submitError, setSubmitError] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const userData = {
      email: formData.get('email'),
      password: formData.get('password'),
    };

    const newErrors = {};
    if (!userData.email) {
      newErrors.email = 'email is required';
    }
    if (!userData.password) {
      newErrors.password = 'Password is required';
    }
    setErrors(newErrors);

    if (newErrors.email || newErrors.password) {
      setSubmitError('Please fill in all required fields.');
      return;
    }

    try {
      loginUser({
        email: userData.email,
        password: userData.password,
      });
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    if (isSuccess && data) {
      dispatch(
        authenticate({
          token: data?.token,
          userid: data?.response.id,
          user_email: data?.response?.email,
        })
      );
      redirect('/dashboard');
    }
  }, [isSuccess, data, dispatch]);

  return (
    <ThemeProvider theme={defaultTheme}>
      <Container component='main' maxWidth='xs'>
        <Typography
          sx={{
            fontWeight: 'bold',
          }}
          component='h5'
          variant='h5'
          align='center'
          gutterBottom={true}
        >
          Sign In
        </Typography>
        {submitError && (
          <Alert severity='error' sx={{ marginBottom: 2 }}>
            {submitError}
          </Alert>
        )}
        {isSuccess && <Alert severity='success'>{data.message}</Alert>}
        {isError && <Alert severity='error'>{error?.data?.error}</Alert>}
        <CssBaseline />
        <Box className='mt-26 flex flex-col items-center'>
          <div
            style={{
              display: 'flex',
              justifyContent: 'center',
              alignItems: 'center',
            }}
          ></div>

          <Box
            component='form'
            noValidate
            onSubmit={handleSubmit}
            className='mt-3'
          >
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  autoComplete='given-name'
                  name='email'
                  required
                  fullWidth
                  id='email'
                  label='Email'
                  autoFocus
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  name='password'
                  label='Password'
                  type='password'
                  id='password'
                  autoComplete='new-password'
                />
              </Grid>
            </Grid>
            <Box
              sx={{
                width: '100%',
                display: 'flex',
                justifyContent: 'center',
                marginTop: 5,
                marginBottom: 5,
              }}
            >
              {isLoading ? (
                <Box sx={{ display: 'flex' }}>
                  <CircularProgress />
                </Box>
              ) : (
                <Button
                  type='submit'
                  sx={{
                    width: '50%',
                  }}
                  style={{ background: '#0170b9' }}
                  variant='contained'
                  className='mt-3 mb-2'
                >
                  Sign In
                </Button>
              )}
            </Box>
            <Box
              sx={{
                width: '100%',
                display: 'flex',
                justifyContent: 'center',
                marginTop: 5,
              }}
            >
              {/* <Grid container justifyContent='center' spacing={2}>
                <Grid item>
                  <Link
                    href='/auth/signup'
                    variant='body2'
                    sx={{
                      textDecoration: 'none',
                    }}
                  >
                    Sign-Up
                  </Link>
                </Grid>
                <Grid item>
                  <Link
                    href='./forgotpassword'
                    variant='body2'
                    sx={{
                      textDecoration: 'none',
                    }}
                  >
                    Forgot password
                  </Link>
                </Grid>
              </Grid> */}
            </Box>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}
