import axios from "axios";
import AsyncStorage from '@react-native-async-storage/async-storage';

// Axios instance for the authentication service
const authApi = axios.create({
  baseURL: 'http://18.230.227.188:8080/api/v1'
});

// Axios instnce for the projects service
const coreApi = axios.create({
    baseURL: 'http://18.230.227.188:8080/api/v1'
});

const projectApi = axios.create({
    baseURL: 'http://18.230.227.188:8080/api/v1/projects'
})

const ceoApi = axios.create({
    baseURL: 'http://18.230.227.188:8080/api/v1/ceo'
});

// Function to store the token in AsyncStorage
const storeToken = async (token) => {
  try {
    await AsyncStorage.setItem('authToken', token);
  } catch (e) {
    console.error('Saving token failed', e);
  }
};

// Function to load the token from AsyncStorage
const getToken = async () => {
  try {
    return await AsyncStorage.getItem('authToken');
  } catch (e) {
    console.error('Loading token failed', e);
    return null;
  }
};

// Interceptors to add the authentication token to all requests
const setupInterceptors = (api) => {
  api.interceptors.request.use(
    async (config) => {
      const token = await getToken();
      console.log("Token loaded:", token);  // Adicione este log para verificar o token
      if (token) {
        config.headers['Authorization'] = `Bearer ${token}`;
        console.log("Authorization Header:", config.headers['Authorization']);  // Adicione este log para verificar o cabeçalho.
      }
      return config;
    },
    error => {
      return Promise.reject(error);
    }
  );
};

// Configure interceptors for both instances
setupInterceptors(authApi);
setupInterceptors(coreApi);
setupInterceptors(projectApi);
setupInterceptors(ceoApi);

// Function to log in using the authentication service
const login = async (email, password) => {
  try {
    const response = await authApi.post('/login', { email, password });
    const { token } = response.data;
    await storeToken(token);
    return token;
  } catch (error) {
    throw error;
  }
};

// Function to register a new user using the authentication service
const register = async (name, email, password, companyName, office, linkedinLink, interest) => {
  try {
    const response = await authApi.post('/register', {
      name,
      email,
      password,
      company_name: companyName,
      office,
      linkedin_link: linkedinLink,
      interest
    });
    console.log('Register response:', response.data);
    return response.data;
  } catch (error) {
    console.error('Registration error:', error);
    throw error;
  }
};

const updateUserInterests = async (userId, interests) => {
  try {
    const response = await authApi.put(`/users/${userId}`, { interest: interests.join(', ') });
    console.log('Interests updated successfully:', response.data);
  } catch (error) {
    console.error('Failed to update interests:', error);
  }
}


export { authApi, coreApi, projectApi, ceoApi, login, storeToken, getToken, register, updateUserInterests };
