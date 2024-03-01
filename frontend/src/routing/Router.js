import React from 'react';
import {
	BrowserRouter as Router,
	Routes,
	Route,
	Navigate,
} from 'react-router-dom';
import Login from '../components/Login';
import App from '../components/Test';
// import Register from './components/Register';
// import Home from './components/Home';
//import ProtectedRoute from './ProtectedRoute';

const AppRouter = () => {
	return (
		<Router>
			<Routes>
				<Route path="/login" element={<Login />} />
				{/* <Route path="/register" element={<Register />} />
				<Route path="/home" element={<ProtectedRoute component={Home} />} /> */}
				<Route path="/test" element={<App />} />
				<Route path="/" element={<Navigate replace to="/login" />} />
				<Route path="*" element={<Navigate replace to="/login" />} />{' '}
				{/* Catch-all route */}
			</Routes>
		</Router>
	);
};

export default AppRouter;
