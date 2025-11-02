<template>
  <div class="register-view">
    <Toast />
    <div class="register-container">
      <Card class="register-card">
        <template #header>
          <div class="card-header">
            <h2>Регистрация</h2>
          </div>
        </template>
        <template #content>
          <form @submit.prevent="handleSubmit" class="register-form">
            <div class="form-field">
              <label for="name">Имя</label>
              <InputText
                id="name"
                v-model="form.name"
                placeholder="Введите имя"
                :class="{ 'p-invalid': errors.name }"
                class="w-full"
              />
              <small v-if="errors.name" class="p-error">{{ errors.name }}</small>
            </div>

            <div class="form-field">
              <label for="email">Email</label>
              <InputText
                id="email"
                v-model="form.email"
                type="email"
                placeholder="Введите email"
                :class="{ 'p-invalid': errors.email }"
                class="w-full"
              />
              <small v-if="errors.email" class="p-error">{{ errors.email }}</small>
            </div>

            <div class="form-field">
              <label for="password">Пароль</label>
              <Password
                id="password"
                v-model="form.password"
                placeholder="Введите пароль"
                :class="{ 'p-invalid': errors.password }"
                toggleMask
                :feedback="true"
                class="w-full"
              />
              <small v-if="errors.password" class="p-error">{{ errors.password }}</small>
            </div>

            <div class="form-field">
              <label for="confirmPassword">Подтверждение пароля</label>
              <Password
                id="confirmPassword"
                v-model="form.confirmPassword"
                placeholder="Повторите пароль"
                :class="{ 'p-invalid': errors.confirmPassword }"
                :feedback="false"
                toggleMask
                class="w-full"
              />
              <small v-if="errors.confirmPassword" class="p-error">{{ errors.confirmPassword }}</small>
            </div>

            <div class="form-field">
              <div class="checkbox-field">
                <Checkbox
                  id="agree"
                  v-model="form.agreeTerms"
                  :binary="true"
                  :class="{ 'p-invalid': errors.agreeTerms }"
                />
                <label for="agree" class="checkbox-label">
                  Я согласен с условиями использования
                </label>
              </div>
              <small v-if="errors.agreeTerms" class="p-error">{{ errors.agreeTerms }}</small>
            </div>

            <div class="form-actions">
              <Button
                type="submit"
                label="Зарегистрироваться"
                icon="pi pi-user-plus"
                :loading="loading"
                class="w-full"
              />
            </div>

            <div class="form-footer">
              <p>
                Уже есть аккаунт?
                <router-link to="/login" class="link">Войти</router-link>
              </p>
            </div>
          </form>
        </template>
      </Card>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'primevue/usetoast'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Button from 'primevue/button'
import Checkbox from 'primevue/checkbox'
import Toast from 'primevue/toast'

const router = useRouter()
const toast = useToast()

const loading = ref(false)
const form = reactive({
  name: '',
  email: '',
  password: '',
  confirmPassword: '',
  agreeTerms: false
})

const errors = reactive({
  name: '',
  email: '',
  password: '',
  confirmPassword: '',
  agreeTerms: ''
})

const validateForm = () => {
  let isValid = true
  errors.name = ''
  errors.email = ''
  errors.password = ''
  errors.confirmPassword = ''
  errors.agreeTerms = ''

  if (!form.name || form.name.trim().length < 2) {
    errors.name = 'Имя должно содержать минимум 2 символа'
    isValid = false
  }

  if (!form.email) {
    errors.email = 'Email обязателен'
    isValid = false
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) {
    errors.email = 'Некорректный email'
    isValid = false
  }

  if (!form.password) {
    errors.password = 'Пароль обязателен'
    isValid = false
  } else if (form.password.length < 6) {
    errors.password = 'Пароль должен содержать минимум 6 символов'
    isValid = false
  }

  if (form.password !== form.confirmPassword) {
    errors.confirmPassword = 'Пароли не совпадают'
    isValid = false
  }

  if (!form.agreeTerms) {
    errors.agreeTerms = 'Необходимо согласие с условиями использования'
    isValid = false
  }

  return isValid
}

const handleSubmit = async () => {
  if (!validateForm()) {
    return
  }

  loading.value = true

  try {
    // TODO: Заменить на реальный API запрос
    // const response = await fetch('/api/auth/register', {
    //   method: 'POST',
    //   headers: { 'Content-Type': 'application/json' },
    //   body: JSON.stringify({
    //     name: form.name,
    //     email: form.email,
    //     password: form.password
    //   })
    // })
    
    // Имитация запроса
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    toast.add({
      severity: 'success',
      summary: 'Успешно',
      detail: 'Регистрация прошла успешно. Вы можете войти в систему.',
      life: 3000
    })

    // Редирект на страницу входа
    setTimeout(() => {
      router.push('/login')
    }, 1500)
    
  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Ошибка',
      detail: error.message || 'Не удалось зарегистрироваться',
      life: 3000
    })
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.register-view {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: calc(100vh - 200px);
  padding: 2rem;
  background: var(--surface-ground);
}

.register-container {
  width: 100%;
  max-width: 500px;
}

.register-card {
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.card-header {
  background: var(--primary-color);
  color: var(--primary-color-text);
  padding: 1.5rem;
  text-align: center;
}

.card-header h2 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
}

.register-form {
  padding: 1.5rem;
}

.form-field {
  margin-bottom: 1.5rem;
}

.form-field label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: var(--text-color);
}

.checkbox-field {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.checkbox-label {
  margin: 0;
  font-weight: normal;
  cursor: pointer;
}

.form-actions {
  margin-top: 2rem;
  margin-bottom: 1rem;
}

.form-footer {
  text-align: center;
  margin-top: 1.5rem;
  padding-top: 1.5rem;
  border-top: 1px solid var(--surface-border);
}

.form-footer p {
  margin: 0;
  color: var(--text-color-secondary);
}

.link {
  color: var(--primary-color);
  text-decoration: none;
  font-weight: 500;
  margin-left: 0.5rem;
}

.link:hover {
  text-decoration: underline;
}

.w-full {
  width: 100%;
}
</style>

