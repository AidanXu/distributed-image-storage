import React, { useCallback } from 'react';
import { useDropzone } from 'react-dropzone';
import {
	AppBar,
	Toolbar,
	Typography,
	Button,
	Container,
	Grid,
	Paper,
} from '@mui/material';
import { styled } from '@mui/system';

// Styled components for custom styling
const CustomAppBar = styled(AppBar)({
	// custom styles
});

const CustomGrid = styled(Grid)({
	// custom styles for the image grid
});

const ImageCard = styled(Paper)({
	// custom styles for the image card
});

const StyledDropzone = styled('div')(({ theme }) => ({
	border: '2px dashed gray',
	borderRadius: '5px',
	padding: '20px',
	textAlign: 'center',
	cursor: 'pointer',
	color: 'gray',
	// You can add more styles here
}));

let images = []; // This should be state or props in a real app

function App() {
	const onDrop = useCallback((acceptedFiles) => {
		// Here, you would handle file upload, including updating state or making an API call
		console.log(acceptedFiles);
	}, []);

	// Initialize the dropzone hook
	const { getRootProps, getInputProps } = useDropzone({ onDrop });

	return (
		<>
			<CustomAppBar position="static">
				<Toolbar>
					<Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
						Image Repository
					</Typography>
					<Button color="inherit">Upload</Button>
					<Button color="inherit">My Images</Button>
				</Toolbar>
			</CustomAppBar>

			<Container>
				<CustomGrid container spacing={2}>
					{/* Map through your images and create a grid item for each */}
					{images.map((image, index) => (
						<Grid item xs={12} sm={6} md={4} key={index}>
							<ImageCard>{/* Image and actions */}</ImageCard>
						</Grid>
					))}
					{/* Your drag-and-drop area */}
					<Grid item xs={12}>
						<StyledDropzone {...getRootProps()}>
							<input {...getInputProps()} />
							<p>Drag 'n' drop some files here, or click to select files</p>
						</StyledDropzone>
					</Grid>
				</CustomGrid>
			</Container>
		</>
	);
}

export default App;
