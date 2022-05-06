import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Layout from './components/Layout';

function App({props}:any) {
  return (
    <Layout>
      <Component {...props}/>
    </Layout>
  );
}

export default App;