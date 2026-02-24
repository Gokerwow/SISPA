<script setup lang="ts">
import { Search, Plus, Bell } from 'lucide-vue-next'
import { ref } from 'vue'

const notifications = ref(3)
</script>

<template>
  <header class="h-20 bg-white/70 backdrop-blur-2xl border-b border-white/20 shadow-sm">
    <div class="h-full px-8 flex items-center justify-between">

      <!-- Search Bar -->
      <div class="flex-1 max-w-2xl">
        <div class="relative group">
          <Search class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 transition-all duration-300 ease-out group-hover:text-purple-500 group-hover:scale-110" />
          <input
            type="search"
            placeholder="Search appointments, customers, services..."
            class="w-full pl-12 pr-4 py-3 bg-slate-50/80 border border-slate-200/60 rounded-2xl text-sm focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all duration-300 ease-out hover:bg-white hover:border-slate-300/60"
          />
        </div>
      </div>

      <!-- Right Side Actions -->
      <div class="flex items-center gap-4 ml-6">

        <!-- New Booking Button -->
        <button class="px-5 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl hover:shadow-purple-300/60 transition-all duration-300 ease-out flex items-center gap-2 hover:scale-105 active:scale-95">
          <Plus class="w-4 h-4 transition-transform duration-300 ease-out group-hover:rotate-90" />
          <span>New Booking</span>
        </button>

        <!-- Notifications -->
        <button class="relative p-3 hover:bg-slate-100/80 rounded-2xl transition-all duration-300 ease-out hover:scale-105 active:scale-95 group">
          <Bell class="w-5 h-5 text-slate-600 transition-all duration-300 ease-out group-hover:text-purple-600 group-hover:rotate-12" />
          <span v-if="notifications > 0" class="absolute top-1.5 right-1.5 w-5 h-5 bg-rose-500 text-white text-xs rounded-full flex items-center justify-center font-bold transition-all duration-300 ease-out group-hover:scale-110">
            {{ notifications }}
          </span>
        </button>

        <!-- User Avatar -->
        <div class="flex items-center gap-3 pl-4 border-l border-slate-200">
          <div class="hidden lg:block text-right">
            <p class="text-sm font-semibold text-slate-700">Olivia Martin</p>
            <p class="text-xs text-slate-500">Manager</p>
          </div>
          <div class="w-11 h-11 rounded-2xl bg-gradient-to-tr from-rose-400 via-purple-400 to-indigo-400 flex items-center justify-center text-white font-bold cursor-pointer hover:scale-110 active:scale-95 transition-all duration-300 ease-out shadow-lg hover:shadow-xl hover:shadow-purple-300/50">
            OM
          </div>
        </div>
      </div>
    </div>
  </header>
</template>
