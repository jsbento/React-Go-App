import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Layout from './components/Layout';
import Pages from './Pages';

function App() {
  return (
    <Layout>
      <Pages/>
    </Layout>
  );
}

export default App;