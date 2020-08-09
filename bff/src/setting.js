import axios from 'axios';
import {constants, _ } from "./constants";

axios.defaults.baseURL = constants.API_SERVER_HOST + ":" + constants.API_SERVER_PORT + "/api/v1";

export default axios;
