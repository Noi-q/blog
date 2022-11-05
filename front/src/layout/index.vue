<script setup lang="ts">
import Header from "./header/index.vue"
import Footer from './footer/index.vue'
</script>

<template>
  <Header/>
  <router-view></router-view>
  <Footer/>
</template>