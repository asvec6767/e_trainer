<template>
  <footer class="app-footer">
    <div class="footer-content">
      <p>&copy; 2025 E-trainer. Все права защищены.</p>
    </div>
  </footer>
</template>

<script setup>
// Компонент Footer
</script>

<style scoped>
.app-footer {
  background: var(--surface-section);
  color: var(--text-color);
  padding: 2rem;
  text-align: center;
  margin-top: auto;
  border-top: 1px solid var(--surface-border);
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
}

.footer-content p {
  margin: 0;
  color: var(--text-color-secondary);
}
</style>

