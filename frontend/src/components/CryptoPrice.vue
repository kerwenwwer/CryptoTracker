<template>
  <div>
    <h1>{{ crypto }} Price: ${{ price }}</h1>
    <button @click="fetchPrice">Refresh Price</button>
  </div>
</template>

<script>
import CryptoService from '../services/CryptoService';

export default {
  data() {
    return {
      price: '',
      crypto: 'BTC'
    };
  },
  methods: {
    fetchPrice() {
      CryptoService.getCryptoPrice(this.crypto)
        .then(response => {
          this.price = response.data.price;
        })
        .catch(error => {
          console.error("There was an error!", error);
        });
    }
  },
  created() {
    this.fetchPrice();
  }
}
</script>
