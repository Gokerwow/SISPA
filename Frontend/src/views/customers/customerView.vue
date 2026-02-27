<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { Users, Plus, Search, Phone, MapPin, Edit2, Trash2, Eye, Filter } from 'lucide-vue-next'
import { RouterLink } from 'vue-router'
import { DeleteCustomers, GetAllCustomers } from '@/services/customers'
import { useToast } from 'vue-toastification'
import type { FullCustomer } from '@/types/customer'

const customers = ref<FullCustomer[]>([])
const loading = ref(true)
const error = ref("")
const searchQuery = ref("")
const toast = useToast()
const refreshTrigger = ref(0)

// 2. Fetch Logic (The "Jarvis" Connection)
const fetchCustomers = async () => {
  try {
    const response = await GetAllCustomers()
    customers.value = response.data
    console.log(customers.value)
  } catch (err) {
    error.value = (err as Error).message
    console.log(error.value)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchCustomers()
})

watch(refreshTrigger, () => {
  fetchCustomers()
})

// Get initials for avatar
const getInitials = (name: string) => {
  return name
    .split(' ')
    .map(n => n[0])
    .join('')
    .toUpperCase()
    .slice(0, 2)
}

// Generate random gradient for avatar
const getAvatarGradient = (id: number) => {
  const gradients = [
    'from-rose-400 to-pink-500',
    'from-purple-400 to-indigo-500',
    'from-blue-400 to-cyan-500',
    'from-emerald-400 to-teal-500',
    'from-orange-400 to-red-500',
    'from-fuchsia-400 to-purple-500'
  ]
  return gradients[id % gradients.length]
}

const handleDelete = async (id: number) => {
  try {
    const result = await DeleteCustomers(id)
    console.log(result)
    toast.success("Customer Deleted Successfully!");
    refreshTrigger.value++
  } catch (error) {
    const message = error.response?.data?.error || "Failed to delete customer";
    console.error('Error deleting customer:', message)
    toast.error(message);
  }
}


</script>

<template>
  <div class="space-y-6">
    <!-- Header Section -->
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
      <div>
        <h1 class="text-3xl font-bold bg-gradient-to-r from-rose-600 via-purple-600 to-indigo-600 bg-clip-text text-transparent mb-2">
          Customers
        </h1>
        <p class="text-slate-500 text-sm">Manage your client relationships</p>
      </div>
      <RouterLink to="/customers/create" class="px-6 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl hover:shadow-purple-300/50 transition-all flex items-center gap-2">
        <Plus class="w-4 h-4" />
        <span>Add New Customer</span>
      </RouterLink>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-purple-400 to-indigo-500 flex items-center justify-center shadow-lg shadow-purple-300/30">
            <Users class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">Total Customers</p>
            <p class="text-2xl font-bold text-slate-800">{{ customers.length }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-emerald-400 to-teal-500 flex items-center justify-center shadow-lg shadow-emerald-300/30">
            <Users class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">Active This Month</p>
            <p class="text-2xl font-bold text-slate-800">{{ Math.floor(customers.length * 0.7) }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-rose-400 to-orange-500 flex items-center justify-center shadow-lg shadow-rose-300/30">
            <Users class="w-6 h-6 text-white" />
          </div>
          <div>
            <p class="text-sm text-slate-500 font-medium">New This Week</p>
            <p class="text-2xl font-bold text-slate-800">{{ Math.floor(customers.length * 0.1) }}</p>
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
            placeholder="Search customers by name, phone, or address..."
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
          <Users class="w-8 h-8 text-white" />
        </div>
        <p class="text-slate-500 font-medium">Loading customers...</p>
      </div>
    </div>

    <!-- Customers List -->
    <div v-else class="bg-white/80 backdrop-blur-xl rounded-2xl border border-white/60 shadow-lg overflow-hidden">
      <!-- Table for larger screens -->
      <div class="hidden lg:block overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-slate-100 bg-slate-50/50">
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Customer</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Contact</th>
              <th class="text-left p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Location</th>
              <th class="text-right p-5 text-xs font-semibold text-slate-500 uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100">
            <tr v-for="customer in customers" :key="customer.id" class="hover:bg-slate-50/50 transition-colors group">
              <td class="p-5">
                <div class="flex items-center gap-3">
                  <div :class="['w-12 h-12 rounded-xl bg-gradient-to-br flex items-center justify-center text-white font-bold shadow-md', getAvatarGradient(customer.id)]">
                    {{ getInitials(customer.name) }}
                  </div>
                  <div>
                    <p class="font-semibold text-slate-800">{{ customer.name }}</p>
                    <p class="text-xs text-slate-500">ID: #{{ customer.id }}</p>
                  </div>
                </div>
              </td>
              <td class="p-5">
                <div class="flex items-center gap-2 text-slate-600">
                  <Phone class="w-4 h-4 text-slate-400" />
                  <span class="text-sm">{{ customer.phone }}</span>
                </div>
              </td>
              <td class="p-5">
                <div class="flex items-center gap-2 text-slate-600">
                  <MapPin class="w-4 h-4 text-slate-400" />
                  <span class="text-sm">{{ customer.address || 'No address' }}</span>
                </div>
              </td>
              <td class="p-5">
                <div class="flex items-center justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                  <button @click="$router.push(`/customers/${customer.id}`)" class="p-2 hover:bg-blue-50 text-blue-600 rounded-lg transition-colors" title="View">
                    <Eye class="w-4 h-4" />
                  </button>
                  <button @click="$router.push(`/customers/edit/${customer.id}`)" class="p-2 hover:bg-purple-50 text-purple-600 rounded-lg transition-colors" title="Edit">
                    <Edit2 class="w-4 h-4" />
                  </button>
                  <button @click="handleDelete(customer.id)" class="p-2 hover:bg-red-50 text-red-600 rounded-lg transition-colors" title="Delete">
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
        <div v-for="customer in customers" :key="customer.id" class="p-5 hover:bg-slate-50/50 transition-colors">
          <div class="flex items-start gap-4">
            <div :class="['w-14 h-14 rounded-xl bg-gradient-to-br flex items-center justify-center text-white font-bold shadow-md flex-shrink-0', getAvatarGradient(customer.id)]">
              {{ getInitials(customer.name) }}
            </div>
            <div class="flex-1 min-w-0">
              <h3 class="font-semibold text-slate-800 mb-1">{{ customer.name }}</h3>
              <div class="space-y-2">
                <div class="flex items-center gap-2 text-slate-600 text-sm">
                  <Phone class="w-3.5 h-3.5 text-slate-400" />
                  <span>{{ customer.phone }}</span>
                </div>
                <div class="flex items-center gap-2 text-slate-600 text-sm">
                  <MapPin class="w-3.5 h-3.5 text-slate-400" />
                  <span>{{ customer.address || 'No address' }}</span>
                </div>
              </div>
              <div class="flex gap-2 mt-3">
                <button class="px-3 py-1.5 bg-blue-50 hover:bg-blue-100 text-blue-600 rounded-lg text-xs font-medium transition-colors flex items-center gap-1">
                  <Eye class="w-3 h-3" />
                  View
                </button>
                <button class="px-3 py-1.5 bg-purple-50 hover:bg-purple-100 text-purple-600 rounded-lg text-xs font-medium transition-colors flex items-center gap-1">
                  <Edit2 class="w-3 h-3" />
                  Edit
                </button>
                <button class="px-3 py-1.5 bg-red-50 hover:bg-red-100 text-red-600 rounded-lg text-xs font-medium transition-colors flex items-center gap-1">
                  <Trash2 class="w-3 h-3" />
                  Delete
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="customers.length === 0" class="p-20 text-center">
        <div class="inline-flex items-center justify-center w-20 h-20 rounded-full bg-gradient-to-r from-slate-100 to-slate-200 mb-4">
          <Users class="w-10 h-10 text-slate-400" />
        </div>
        <h3 class="text-lg font-semibold text-slate-700 mb-2">No customers yet</h3>
        <p class="text-slate-500 text-sm mb-6">Start building your client base by adding your first customer</p>
        <button class="px-6 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl transition-all inline-flex items-center gap-2">
          <Plus class="w-4 h-4" />
          <span>Add Your First Customer</span>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Smooth transitions */
* {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Smooth font rendering */
body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
