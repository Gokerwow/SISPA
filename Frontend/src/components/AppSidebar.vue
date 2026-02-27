<script setup lang="ts">
import { RouterLink } from 'vue-router'
import {
  LayoutDashboard,
  CalendarDays,
  Users,
  Heart,
  UserCircle,
  Package,
  TrendingUp,
  Settings,
  Sparkles,
  ChevronRight,
  LogOut,
  FileText
} from 'lucide-vue-next'

const menuItems = [
  { path: '/dashboard', icon: LayoutDashboard, label: 'Dashboard' },
  { path: '/appointments', icon: CalendarDays, label: 'Appointments' },
  { path: '/transactions', icon: FileText, label: 'Transactions' },
  { path: '/customers', icon: Users, label: 'Customers' },
  { path: '/services', icon: Heart, label: 'Services' },
  { path: '/therapists', icon: UserCircle, label: 'Therapists' },
  { path: '/inventory', icon: Package, label: 'Inventory' }
]

const analyticsItems = [
  { path: '/reports', icon: TrendingUp, label: 'Reports' },
  { path: '/settings', icon: Settings, label: 'Settings' }
]
</script>

<template>
  <aside class="w-72 bg-white/80 backdrop-blur-2xl border-r border-white/60 shadow-xl flex flex-col">

    <!-- Logo & Brand -->
    <div class="p-6 border-b border-slate-100">
      <div class="flex items-center gap-3">
        <div class="relative group">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-tr from-rose-400 via-purple-400 to-indigo-400 flex items-center justify-center shadow-xl shadow-purple-300/50 rotate-6 group-hover:rotate-0 transition-transform duration-500 ease-out">
            <Sparkles class="w-6 h-6 text-white" :stroke-width="2.5" />
          </div>
          <div class="absolute -bottom-1 -right-1 w-4 h-4 bg-green-400 rounded-full border-2 border-white"></div>
        </div>
        <div>
          <h1 class="text-2xl font-bold bg-gradient-to-r from-rose-600 via-purple-600 to-indigo-600 bg-clip-text text-transparent">
            SISPA
          </h1>
          <p class="text-xs text-slate-500 font-medium -mt-1">Spa Excellence</p>
        </div>
      </div>
    </div>

    <!-- Navigation Menu -->
    <nav class="flex-1 p-4 space-y-2 overflow-y-auto">
      <!-- Main Menu Items -->
      <RouterLink
        v-for="item in menuItems"
        :key="item.path"
        :to="item.path"
        class="flex items-center gap-3 px-4 py-3.5 rounded-2xl text-slate-700 transition-all duration-300 ease-out group relative overflow-hidden"
        active-class="bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white shadow-lg shadow-purple-300/40">

        <!-- Hover Background -->
        <span class="absolute inset-0 bg-gradient-to-r from-rose-50 to-purple-50 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity duration-300 ease-out -z-10"></span>

        <component :is="item.icon" class="w-5 h-5 relative z-10 transition-transform duration-300 ease-out group-hover:scale-110" :stroke-width="2" />
        <span class="font-medium text-sm flex-1 relative z-10">{{ item.label }}</span>
        <ChevronRight class="w-4 h-4 opacity-0 group-hover:opacity-100 transition-all duration-300 ease-out group-hover:translate-x-1" />
      </RouterLink>

      <!-- Section Divider -->
      <div class="pt-4 pb-2">
        <p class="px-4 text-xs font-semibold text-slate-400 uppercase tracking-wider">Analytics</p>
      </div>

      <!-- Analytics Items -->
      <RouterLink
        v-for="item in analyticsItems"
        :key="item.path"
        :to="item.path"
        class="flex items-center gap-3 px-4 py-3.5 rounded-2xl text-slate-700 transition-all duration-300 ease-out group relative overflow-hidden"
        active-class="bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white shadow-lg shadow-purple-300/40">

        <!-- Hover Background -->
        <span class="absolute inset-0 bg-gradient-to-r from-rose-50 to-purple-50 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity duration-300 ease-out -z-10"></span>

        <component :is="item.icon" class="w-5 h-5 relative z-10 transition-transform duration-300 ease-out group-hover:scale-110" :stroke-width="2" />
        <span class="font-medium text-sm flex-1 relative z-10">{{ item.label }}</span>
        <ChevronRight class="w-4 h-4 opacity-0 group-hover:opacity-100 transition-all duration-300 ease-out group-hover:translate-x-1" />
      </RouterLink>
    </nav>

    <!-- User Profile Section -->
    <div class="p-4 border-t border-slate-100">
      <div class="p-4 rounded-2xl bg-gradient-to-br from-slate-50 to-purple-50 border border-slate-200/60 transition-all duration-300 ease-out hover:shadow-lg hover:shadow-purple-200/50">
        <div class="flex items-center gap-3 mb-3">
          <div class="w-11 h-11 rounded-2xl bg-gradient-to-tr from-rose-400 via-purple-400 to-indigo-400 flex items-center justify-center text-white font-bold shadow-lg transition-transform duration-300 ease-out hover:scale-105">
            OM
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm font-semibold text-slate-800 truncate">Olivia Martin</p>
            <p class="text-xs text-slate-500">Spa Manager</p>
          </div>
        </div>
        <button class="w-full px-4 py-2.5 bg-white hover:bg-red-50 text-red-600 rounded-xl text-sm font-medium transition-all duration-300 ease-out flex items-center justify-center gap-2 border border-red-200/60 hover:border-red-300 hover:shadow-md">
          <LogOut class="w-4 h-4 transition-transform duration-300 ease-out group-hover:translate-x-1" />
          Logout
        </button>
      </div>
    </div>
  </aside>
</template>

<style scoped>
/* Custom scrollbar */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #e879f9, #a78bfa);
  border-radius: 10px;
}

::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #d946ef, #8b5cf6);
}
</style>
