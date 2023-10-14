'use client';

import React, { useEffect, useState } from 'react';
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from '@mui/material';
import { useGetAllLoansQuery } from '@/slices/loan-slice';

const LoansList = () => {
  const [loanData, setLoanData] = useState([]);
  const { isLoading, isSuccess, isError, error, data } = useGetAllLoansQuery();

  useEffect(() => {
    if (isSuccess && data) {
      setLoanData(data.response);
    }
  }, [isSuccess && data]);
  return (
    <TableContainer component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Customer ID</TableCell>
            <TableCell>Loan Amount</TableCell>
            <TableCell>Paid Amount</TableCell>
            <TableCell>Interest</TableCell>
            <TableCell>Principal</TableCell>
            <TableCell>Loan Time</TableCell>
            <TableCell>Amount to be Paid</TableCell>
            <TableCell>Rate</TableCell>
            <TableCell>Due Date</TableCell>
            <TableCell>Remaining Amount</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {loanData?.map((row, index) => (
            <TableRow key={index}>
              <TableCell>{row.customer_id}</TableCell>
              <TableCell>{row.loan_amount}</TableCell>
              <TableCell>{row.paid_amount}</TableCell>
              <TableCell>{row.interest}</TableCell>
              <TableCell>{row.principal}</TableCell>
              <TableCell>{row.loan_time}</TableCell>
              <TableCell>{row.amount_to_be_paid}</TableCell>
              <TableCell>{row.rate}</TableCell>
              <TableCell>{row.due_date}</TableCell>
              <TableCell>{row.remaining_amount}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default LoansList;
