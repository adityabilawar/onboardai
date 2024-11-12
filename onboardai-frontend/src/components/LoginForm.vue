<template>
    <form @submit.prevent="handleSubmit" class="w-full max-w-lg mx-auto p-6">
      <div class="text-center mb-8">
        <h1 class="text-2xl font-bold text-gray-900">Onboard AI</h1>
      </div>
      <div v-if="userStore.error" class="mb-4 p-3 bg-red-100 text-red-700 rounded">
        {{ userStore.error }}
      </div>
  
      <div class="mb-4">
        <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email</label>
        <input
          id="email"
          v-model="email"
          type="email"
          required
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="{ 'border-red-500': errors.email }"
          placeholder="Enter your email"
          :disabled="userStore.isLoading"
        >
        <p v-if="errors.email" class="mt-1 text-sm text-red-500">{{ errors.email }}</p>
      </div>
  
      <div class="mb-6">
        <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
        <input
          id="password"
          v-model="password"
          type="password"
          required
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="{ 'border-red-500': errors.password }"
          placeholder="Enter your password"
          :disabled="userStore.isLoading"
        >
        <p v-if="errors.password" class="mt-1 text-sm text-red-500">{{ errors.password }}</p>
      </div>
  
      <div class="flex justify-between items-center">
        <button
          type="submit"
          class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50"
          :disabled="userStore.isLoading"
        >
          <span v-if="userStore.isLoading">
            Processing...
          </span>
          <span v-else>
            {{ isLogin ? 'Login' : 'Sign Up' }}
          </span>
        </button>
        <button
          type="button"
          @click="toggleMode"
          class="text-blue-500 hover:text-blue-600 text-sm"
          :disabled="userStore.isLoading"
        >
          {{ isLogin ? 'Need an account? Sign up' : 'Already have an account? Login' }}
        </button>
      </div>
    </form>
  </template>
  
  <script setup>
  import { ref, reactive } from 'vue'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '../stores/user'
  
  const router = useRouter()
  const userStore = useUserStore()
  
  const email = ref('')
  const password = ref('')
  const isLogin = ref(true)
  const errors = reactive({
    email: '',
    password: ''
  })
  
  const validateForm = () => {
    let isValid = true
    errors.email = ''
    errors.password = ''
  
    if (!email.value) {
      errors.email = 'Email is required'
      isValid = false
    }
    if (!password.value) {
      errors.password = 'Password is required'
      isValid = false
    } else if (password.value.length < 6) {
      errors.password = 'Password must be at least 6 characters'
      isValid = false
    }
  
    return isValid
  }
  
  const handleSubmit = async () => {
    if (!validateForm()) return
  
    const success = isLogin.value 
      ? await userStore.login(email.value, password.value)
      : await userStore.signup(email.value, password.value)
  
    if (success) {
      const role = userStore.userRole
      if (role === 'manager') {
        router.push('/manager-dashboard')
      } else if (role === 'engineer') {
        router.push('/engineer-dashboard')
      } else {
        router.push('/role-selection')
      }
    }
  }
  
  const toggleMode = () => {
    isLogin.value = !isLogin.value
    errors.email = ''
    errors.password = ''
    userStore.error = null
  }
  </script>