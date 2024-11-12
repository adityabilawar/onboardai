import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')) || null,
    role: localStorage.getItem('userRole') || null,
    isLoading: false,
    error: null
  }),

  actions: {
    async login(email, password) {
      this.isLoading = true
      this.error = null
      
      try {
        const response = await fetch('http://localhost:8080/api/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email, password })
        })

        if (!response.ok) {
          throw new Error('Invalid credentials')
        }

        const data = await response.json()
        const user = {
          email: data.user.email,
          role: data.user.role
        }

        localStorage.setItem('user', JSON.stringify(user))
        this.user = user
        return true
      } catch (e) {
        this.error = e.message
        return false
      } finally {
        this.isLoading = false
      }
    },

    async signup(email, password) {
      this.isLoading = true
      this.error = null
      
      try {
        const response = await fetch('http://localhost:8080/api/signup', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email, password })
        })

        if (!response.ok) {
          throw new Error('Signup failed')
        }

        const data = await response.json()
        const user = {
          email: data.user.email,
          role: data.user.role
        }

        localStorage.setItem('user', JSON.stringify(user))
        this.user = user
        return true
      } catch (e) {
        this.error = e.message
        return false
      } finally {
        this.isLoading = false
      }
    },

    setRole(role) {
      this.role = role
      localStorage.setItem('userRole', role)
      
      // Update user object with role
      if (this.user) {
        this.user = { ...this.user, role }
        localStorage.setItem('user', JSON.stringify(this.user))
      }
    },

    logout() {
      this.user = null
      this.role = null
      this.error = null
      localStorage.clear()
    }
  },

  getters: {
    isAuthenticated: (state) => !!state.user,
    userEmail: (state) => state.user?.email,
    userRole: (state) => state.user?.role || state.role
  }
})