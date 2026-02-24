<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppSidebar from '@/components/AppSidebar.vue';
import AppHeader from '@/components/AppHeader.vue'
</script>

<template>
  <div class="flex h-screen bg-gradient-to-br from-rose-50 via-purple-50 to-indigo-50">

    <!-- Sidebar Component -->
    <AppSidebar />

    <!-- Main Content Area -->
    <div class="flex-1 flex flex-col overflow-hidden">

      <!-- Header Component -->
      <AppHeader />

      <!-- Page Content -->
      <main class="flex-1 overflow-auto p-8">
        <RouterView />
      </main>
    </div>
  </div>
</template>

<style>
/* Global smooth transitions */
* {
  transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 300ms;
}

/* Custom scrollbar for main content */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #e879f9, #a78bfa);
  border-radius: 10px;
  border: 2px solid transparent;
  background-clip: padding-box;
}

::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #d946ef, #8b5cf6);
  background-clip: padding-box;
}

/* Smooth font rendering */
body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
