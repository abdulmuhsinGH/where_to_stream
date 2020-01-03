import axios from 'axios';

const link = axios.create({
  baseURL: `${process.env.VUE_APP_API_URL}`,
});

export default link;
