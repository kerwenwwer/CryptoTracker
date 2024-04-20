import axios from 'axios';

const API_URL = 'http://localhost:8080/price/';

class CryptoService {
    getCryptoPrice(crypto) {
        return axios.get(`${API_URL}${crypto}`);
    }
}

export default new CryptoService();
