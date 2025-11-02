<template>
  <div class="home-view">
    <Toast />
    <div class="container">
      <Card>
        <template #header>
          <div class="card-header">
            <h1>E-trainer</h1>
          </div>
        </template>
        <template #content>
          <div class="content">
            <p>Добро пожаловать в E-trainer!</p>
            <Button label="Начать" icon="pi pi-check" @click="handleClick" />
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>

<script setup>
import Card from 'primevue/card'
import Button from 'primevue/button'
import Toast from 'primevue/toast'
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const handleClick = () => {
  toast.add({ 
    severity: 'info', 
    summary: 'Успешно', 
    detail: 'Приложение работает!', 
    life: 3000 
  })
}
</script>

<style scoped>
.home-view {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  min-height: calc(100vh - 200px);
  background: var(--surface-ground);
}

.container {
  width: 100%;
  max-width: 600px;
}

.card-header {
  background: var(--primary-color);
  color: var(--primary-color-text);
  padding: 2rem;
  text-align: center;
}

.card-header h1 {
  margin: 0;
  font-size: 2rem;
}

.content {
  padding: 2rem;
  text-align: center;
}

.content p {
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
  color: var(--text-color);
}
</style>

