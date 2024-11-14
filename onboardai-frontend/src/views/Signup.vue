<!-- src/views/Login.vue -->
<template>
    <div class="login">
      <h2>Sign Up</h2>
      <form @submit.prevent="handleSubmit">
        <input 
          v-model="email" 
          type="email" 
          placeholder="Email" 
          required
        >
        <input 
          v-model="password" 
          type="password" 
          placeholder="Password" 
          required
        >
        <button type="submit">Submit</button>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '../stores/user'
  
  const router = useRouter()
  const userStore = useUserStore()
  
  const email = ref('')
  const password = ref('')
  
  const handleSubmit = async () => {
    await userStore.signup(email.value, password.value)
    router.push('/role-selection')
  }
  </script>
  