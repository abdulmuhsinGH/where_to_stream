import axios from 'axios';

// axios.defaults.headers.common['Access-Control-Allow-Origin'] = `${process.env.VUE_APP_API_URL}`;
const link = axios.create({
  baseURL: '/',
});

export default link;
