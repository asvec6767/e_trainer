<template>
  <header class="app-header">
    <div class="header-content">
      <router-link to="/" class="logo">
        <h1>E-trainer</h1>
      </router-link>
      <nav class="navigation">
        <router-link to="/" class="nav-link">Главная</router-link>
        <router-link to="/login" class="nav-link" v-if="!isAuthenticated">
          <i class="pi pi-sign-in"></i>
          Войти
        </router-link>
        <router-link to="/register" class="nav-link register-btn" v-if="!isAuthenticated">
          <i class="pi pi-user-plus"></i>
          Регистрация
        </router-link>
        <!-- TODO: Добавить меню для авторизованного пользователя -->
      </nav>
    </div>
  </header>
</template>

<script setup>
import { ref } from 'vue'

// TODO: Заменить на реальную проверку авторизации из store/API
const isAuthenticated = ref(false)
</script>

<style scoped>
.app-header {
  background: var(--primary-color);
  color: var(--primary-color-text);
  padding: 1rem 2rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  text-decoration: none;
  color: inherit;
  cursor: pointer;
}

.logo:hover {
  opacity: 0.9;
}

.logo h1 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
}

.navigation {
  display: flex;
  gap: 1.5rem;
}

.nav-link {
  color: var(--primary-color-text);
  text-decoration: none;
  font-weight: 500;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: background-color 0.2s;
}

.nav-link:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.nav-link.router-link-active {
  background-color: rgba(255, 255, 255, 0.2);
}

.nav-link i {
  margin-right: 0.5rem;
}

.register-btn {
  background-color: rgba(255, 255, 255, 0.15);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.register-btn:hover {
  background-color: rgba(255, 255, 255, 0.25);
}
</style>

