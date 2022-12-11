import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Link from '@mui/material/Link';
import './App.css';


import { BrowserRouter, Route} from 'react-router-dom';
import Users from './pages/Users';
import Login from './pages/Login';
import Register from './pages/Register';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Route path={'/'} exact component={Users}></Route>
        <Route path={'/login'} component={Login}></Route>
        <Route path={'/register'} component={Register}></Route>
      </BrowserRouter>  
    </div>
  );
}
export default App;
