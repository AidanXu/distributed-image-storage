import React from 'react';
import { Route, Navigate } from 'react-router-dom';

const ProtectedRoute = ({ component: Component, ...rest }) => {
	const token = localStorage.getItem('token');
	const isAuthenticated = token ? true : false;
	return (
		<Route
			{...rest}
			render={(props) =>
				isAuthenticated ? <Component {...props} /> : <Navigate to="/login" />
			}
		/>
	);
};

export default ProtectedRoute;
