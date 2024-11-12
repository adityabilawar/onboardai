// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import RoleSelection from '../views/RoleSelection.vue'
import ManagerDashboard from '../views/ManagerDashboard.vue'
import EngineerDashboard from '../views/EngineerDashboard.vue'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
  {
    path: '/role-selection',
    name: 'RoleSelection',
    component: RoleSelection,
    meta: { requiresAuth: true }
  },
  {
    path: '/manager-dashboard',
    name: 'ManagerDashboard',
    component: ManagerDashboard,
    meta: { requiresAuth: true }
  },
  {
    path: '/engineer-dashboard',
    name: 'EngineerDashboard',
    component: EngineerDashboard,
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation guards
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('user')
  const user = isAuthenticated ? JSON.parse(localStorage.getItem('user')) : null
  
  if (to.path === '/' && isAuthenticated && user?.role) {
    // Redirect logged in users with roles to their dashboard
    next(user.role === 'manager' ? '/manager-dashboard' : '/engineer-dashboard')
  } else if (to.meta.requiresAuth && !isAuthenticated) {
    next('/')
  } else {
    next()
  }
})

export default router