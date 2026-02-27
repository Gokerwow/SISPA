<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { Heart, Plus, Search, DollarSign, Edit2, Trash2, Eye, Filter, Sparkles } from 'lucide-vue-next'
import { RouterLink } from 'vue-router'
import { GetAllServices } from '@/services/services'

// Service interface
export interface Services {
  id: number
  name: string
  price_male: number
  price_female: number
  price_tourist: number
  is_active: boolean
  created_at: string
}

const services = ref<Services[]>([])
const loading = ref(true)
const error = ref("")
const searchQuery = ref("")
const refreshTrigger = ref(0)

// Fetch Logic
const fetchServices = async () => {
  try {
    const response = await GetAllServices()
    services.value = response.data
    console.log(services.value)
  } catch (err) {
    error.value = (err as Error).message
    console.log(error.value)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchServices()
})

watch(refreshTrigger, () => {
  fetchServices()
})

// Format currency
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount)
}

// Get service icon based on name
const getServiceIcon = (name: string) => {
  const nameLower = name.toLowerCase()
  if (nameLower.includes('massage')) return '💆‍♀️'
  if (nameLower.includes('facial')) return '✨'
  if (nameLower.includes('body')) return '🧖‍♀️'
  if (nameLower.includes('nail') || nameLower.includes('manicure') || nameLower.includes('pedicure')) return '💅'
  if (nameLower.includes('hair')) return '💇‍♀️'
  if (nameLower.includes('spa')) return '🌸'
  return '💎'
}

// Get service gradient based on ID
const getServiceGradient = (id: number) => {
  const gradients = [
    'from-purple-400 to-indigo-500',
    'from-rose-400 to-pink-500',
    'from-emerald-400 to-teal-500',
    'from-orange-400 to-red-500',
    'from-blue-400 to-cyan-500',
    'from-fuchsia-400 to-purple-500'
  ]
  return gradients[id % gradients.length]
}

const handleDelete = async (id: number) => {
  try {
    const response = await fetch(`http://localhost:8080/services/${id}`, {
      method: 'DELETE'
    })
    if (!response.ok) throw new Error("Failed to delete service")

    console.log('Service deleted')
    // Show success toast here
    refreshTrigger.value++
  } catch (error) {
    console.error('Error deleting service:', error)
    // Show error toast here
  }
}

// Calculate stats
const activeServices = ref(0)
const averagePrice = ref(0)

watch(services, (newServices) => {
  activeServices.value = newServices.filter(s => s.is_active).length
  // Calculate average across all price types
  if (newServices.length > 0) {
    const totalPrices = newServices.reduce((sum, s) => {
      return sum + s.price_male + s.price_female + s.price_tourist
    }, 0)
    averagePrice.value = totalPrices / (newServices.length * 3)
  } else {
    averagePrice.value = 0
  }
}, { immediate: true })
</script>

<template>
  <div class="space-y-6">
    <!-- Header Section -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-3xl font-bold bg-gradient-to-r from-rose-600 via-purple-600 to-indigo-600 bg-clip-text text-transparent mb-2">
          Services
        </h1>
        <p class="text-slate-500 text-sm">Manage your spa treatments and offerings</p>
      </div>
      <RouterLink to="/services/create" class="px-6 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl hover:shadow-purple-300/50 transition-all flex items-center gap-2">
        <Plus class="w-4 h-4" />
        <span>Add New Service</span>
      </RouterLink>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-purple-400 to-indigo-500 flex items-center justify-center shadow-lg shadow-purple-300/30">
            <Heart class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">Total Services</p>
            <p class="text-2xl font-bold text-slate-800">{{ services.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-emerald-400 to-teal-500 flex items-center justify-center shadow-lg shadow-emerald-300/30">
            <Sparkles class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">Active Services</p>
            <p class="text-2xl font-bold text-slate-800">{{ activeServices }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-rose-400 to-orange-500 flex items-center justify-center shadow-lg shadow-rose-300/30">
            <DollarSign class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">Average Price</p>
            <p class="text-2xl font-bold text-slate-800">{{ formatCurrency(averagePrice) }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Search & Filter Bar -->
    <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-4 border border-white/60 shadow-lg">
      <div class="flex flex-col sm:flex-row gap-3">
        <div class="flex-1 relative">
          <Search class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
          <input
            v-model="searchQuery"
            type="search"
            placeholder="Search services by name..."
            class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all"
          />
        </div>
        <button class="px-5 py-3 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-xl text-sm font-medium transition-all flex items-center justify-center gap-2">
          <Filter class="w-4 h-4" />
          <span>Filters</span>
        </button>
      </div>
    </div>

    <!-- Error Message -->
    <div v-if="error" class="bg-red-50 border-l-4 border-red-500 rounded-xl p-5 shadow-lg">
      <div class="flex items-start gap-3">
        <div class="w-10 h-10 rounded-xl bg-red-100 flex items-center justify-center flex-shrink-0">
          <span class="text-red-600 text-xl">⚠️</span>
        </div>
        <div>
          <p class="font-bold text-red-800 mb-1">Connection Error</p>
          <p class="text-red-700 text-sm">{{ error }}</p>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="bg-white/80 backdrop-blur-xl rounded-2xl border border-white/60 shadow-lg p-20">
      <div class="text-center">
        <div class="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gradient-to-r from-rose-400 via-purple-400 to-indigo-400 mb-4 animate-pulse">
          <Heart class="w-8 h-8 text-white" />
        </div>
        <p class="text-slate-500 font-medium">Loading services...</p>
      </div>
    </div>

    <!-- Services List -->
    <div v-else class="bg-white/80 backdrop-blur-xl rounded-2xl border border-white/60 shadow-lg overflow-hidden">
      <!-- Table for larger screens -->
      <div class="hidden lg:block overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-slate-100 bg-slate-50/50">
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Service</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Price (Male)</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Price (Female)</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Price (Tourist)</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Status</th>
              <th class="text-right p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100">
            <tr v-for="service in services" :key="service.id" class="hover:bg-slate-50/50 transition-colors group">
              <td class="p-5">
                <div class="flex items-center gap-3">
                  <div :class="['w-12 h-12 rounded-xl bg-gradient-to-br flex items-center justify-center text-2xl shadow-md', getServiceGradient(service.id)]">
                    {{ getServiceIcon(service.name) }}
                  </div>
                  <div>
                    <p class="font-semibold text-slate-800">{{ service.name }}</p>
                    <p class="text-xs text-slate-500">ID: #{{ service.id }}</p>
                  </div>
                </div>
              </td>
              <td class="p-5">
                <div class="flex items-center gap-1 text-slate-800 font-semibold">
                  <DollarSign class="w-4 h-4 text-blue-400" />
                  <span class="text-sm">{{ formatCurrency(service.price_male) }}</span>
                </div>
              </td>
              <td class="p-5">
                <div class="flex items-center gap-1 text-slate-800 font-semibold">
                  <DollarSign class="w-4 h-4 text-pink-400" />
                  <span class="text-sm">{{ formatCurrency(service.price_female) }}</span>
                </div>
              </td>
              <td class="p-5">
                <div class="flex items-center gap-1 text-slate-800 font-semibold">
                  <DollarSign class="w-4 h-4 text-purple-400" />
                  <span class="text-sm">{{ formatCurrency(service.price_tourist) }}</span>
                </div>
              </td>
              <td class="p-5">
                <span v-if="service.is_active" class="px-3 py-1 bg-emerald-100 text-emerald-700 rounded-lg text-xs font-bold">
                  Active
                </span>
                <span v-else class="px-3 py-1 bg-slate-100 text-slate-500 rounded-lg text-xs font-bold">
                  Inactive
                </span>
              </td>
              <td class="p-5">
                <div class="flex items-center justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                  <button @click="$router.push(`/services/edit/${service.id}`)" class="p-2 hover:bg-purple-50 text-purple-600 rounded-lg transition-colors" title="Edit">
                    <Edit2 class="w-4 h-4" />
                  </button>
                  <button @click="handleDelete(service.id)" class="p-2 hover:bg-red-50 text-red-600 rounded-lg transition-colors" title="Delete">
                    <Trash2 class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Card view for mobile -->
      <div class="lg:hidden divide-y divide-slate-100">
        <div v-for="service in services" :key="service.id" class="p-5 hover:bg-slate-50/50 transition-colors">
          <div class="flex items-start gap-4">
            <div :class="['w-14 h-14 rounded-xl bg-gradient-to-br flex items-center justify-center text-2xl shadow-md flex-shrink-0', getServiceGradient(service.id)]">
              {{ getServiceIcon(service.name) }}
            </div>
            <div class="flex-1 min-w-0">
              <div class="flex items-start justify-between mb-2">
                <h3 class="font-semibold text-slate-800 mb-1">{{ service.name }}</h3>
                <span v-if="service.is_active" class="px-2 py-1 bg-emerald-100 text-emerald-700 rounded-lg text-xs font-bold">
                  Active
                </span>
              </div>
              <div class="space-y-2 mb-3 text-sm">
                <div class="flex items-center justify-between p-2 bg-blue-50 rounded-lg">
                  <span class="text-slate-600 text-xs">Male</span>
                  <span class="font-bold text-slate-800">{{ formatCurrency(service.price_male) }}</span>
                </div>
                <div class="flex items-center justify-between p-2 bg-pink-50 rounded-lg">
                  <span class="text-slate-600 text-xs">Female</span>
                  <span class="font-bold text-slate-800">{{ formatCurrency(service.price_female) }}</span>
                </div>
                <div class="flex items-center justify-between p-2 bg-purple-50 rounded-lg">
                  <span class="text-slate-600 text-xs">Tourist</span>
                  <span class="font-bold text-slate-800">{{ formatCurrency(service.price_tourist) }}</span>
                </div>
              </div>
              <div class="flex gap-2">
                <button @click="$router.push(`/services/${service.id}`)" class="px-3 py-1.5 bg-blue-50 hover:bg-blue-100 text-blue-600 rounded-lg text-xs font-medium transition-colors flex items-center gap-1">
                  <Eye class="w-3 h-3" />
                  View
                </button>
                <button @click="$router.push(`/services/edit/${service.id}`)" class="px-3 py-1.5 bg-purple-50 hover:bg-purple-100 text-purple-600 rounded-lg text-xs font-medium transition-colors flex items-center gap-1">
                  <Edit2 class="w-3 h-3" />
                  Edit
                </button>
                <button @click="handleDelete(service.id)" class="px-3 py-1.5 bg-red-50 hover:bg-red-100 text-red-600 rounded-lg text-xs font-medium transition-colors flex items-center gap-1">
                  <Trash2 class="w-3 h-3" />
                  Delete
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="services.length === 0" class="p-20 text-center">
        <div class="inline-flex items-center justify-center w-20 h-20 rounded-full bg-gradient-to-r from-slate-100 to-slate-200 mb-4">
          <Heart class="w-10 h-10 text-slate-400" />
        </div>
        <h3 class="text-lg font-semibold text-slate-700 mb-2">No services yet</h3>
        <p class="text-slate-500 text-sm mb-6">Start by adding your first spa service</p>
        <RouterLink to="/services/create" class="px-6 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl transition-all inline-flex items-center gap-2">
          <Plus class="w-4 h-4" />
          <span>Add Your First Service</span>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Smooth transitions */
* {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Line clamp for descriptions */
.line-clamp-1 {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Smooth font rendering */
body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
