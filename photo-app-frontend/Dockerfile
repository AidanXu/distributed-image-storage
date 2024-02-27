# Use an official Node runtime as a parent image
FROM node:latest

# Set the working directory in the container
WORKDIR /app

# Copy package.json and the package-lock.json files
COPY package.json package-lock.json ./

# Install any dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Make port 3000 available outside this container
EXPOSE 3000

# Run the app when the container launches
CMD ["npm", "start"]
