import React, { useState, useEffect } from 'react';

function ImageGallery() {
	let images = [];

	useEffect(() => {
		const fetchImages = async () => {
			const jwtToken =
				'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ4ZGQifQ._2GusP9eNaM34zN_Pne78jrETtuv2clrw9iayGfhfKk'; // Replace this with the actual method of obtaining your JWT
			try {
				const response = await fetch('http://localhost:8080/storage', {
					method: 'GET',
					headers: {
						Authorization: `Bearer ${jwtToken}`,
					},
				});
				if (!response.ok) {
					throw new Error(`Error: ${response.status}`);
				}
				const data = await response.json();
				images = data.photos;
				console.log(images);
			} catch (error) {
				console.error('Failed to fetch images:', error);
			}
		};

		fetchImages();
	}, []); // Empty dependency array means this effect runs once on mount

	return (
		<div>
			<h2>Image Gallery</h2>
			<div>
				<img src={images[0]} style={{ width: '100px', height: '100px' }} />
			</div>
		</div>
	);
}

export default ImageGallery;
