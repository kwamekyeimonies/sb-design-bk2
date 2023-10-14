import { createSlice } from '@reduxjs/toolkit';
import Cookies from 'js-cookie';

const authenticationReducer = createSlice({
  name: 'auth',
  initialState: {
    user_email: null,
    username: null,
    token: null,
    userid: null,
    loading: false,
  },
  reducers: {
    authenticate: (state, action) => {
      if (typeof window !== 'undefined') {
        const { user_email, token, userid, username } = action.payload;
        state.user_email = user_email;
        state.token = token;
        state.userid = userid;
        state.username = username;
        state.loading = false;
        localStorage.setItem('user_email', JSON.stringify(user_email));
        localStorage.setItem('token', JSON.stringify(token));
        localStorage.setItem('id', JSON.stringify(userid));
        sessionStorage.setItem('username', username);
        Cookies.set('user_email', user_email, {
          expires: 7,
          httpOnly: false,
          secure: false,
        });
        Cookies.set('username', username, {
          expires: 7,
          httpOnly: false,
          secure: false,
        });
        Cookies.set('token', token, {
          expires: 7,
          httpOnly: false,
          secure: false,
        });
        Cookies.set('id', userid, {
          expires: 7,
          httpOnly: false,
          secure: false,
        });
      }
    },
    startAuthentication: (state) => {
      state.loading = true;
    },
    logout: (state) => {
      if (typeof window !== 'undefined') {
        state.user_email = null;
        state.token = null;
        localStorage.removeItem('user_email');
        localStorage.removeItem('token');
        Cookies.remove('user_email');
        Cookies.remove('token');
      }
    },
  },
});

export const { authenticate, logout, startAuthentication } =
  authenticationReducer.actions;
export default authenticationReducer.reducer;

export const selectCurrentUser = (state) => state.auth.user;
export const selectCurrentToken = (state) => state.auth.token;
