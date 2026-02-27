<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import {
  User,
  Phone,
  MapPin,
  Calendar,
  DollarSign,
  TrendingUp,
  Clock,
  Edit2,
  Trash2,
  ArrowLeft,
  Gift,
  Star,
  Activity,
  FileText
} from 'lucide-vue-next'
import { DeleteCustomers, GetCustomerDetail } from '@/services/customers'
import { useRoute, useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import type { FullCustomer } from '@/types/customer'

const route = useRoute()

const customer = ref<FullCustomer | null>(null)
const loading = ref(true)
const error = ref('')
const router = useRouter()
const toast = useToast()
const id = Number(route.params.id)

// Fetch customer data
const fetchCustomer = async (customerId: number) => {
  try {
    const response = await GetCustomerDetail(customerId)
    customer.value = await response.data
    console.log(response)
  } catch (err) {
    error.value = (err as Error).message
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  if (id) {
    fetchCustomer(id)
  }
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

// Format currency
const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD'
  }).format(amount)
}

// Format date
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

// Calculate age
const calculateAge = (birthDate: string) => {
  const today = new Date()
  const birth = new Date(birthDate)
  let age = today.getFullYear() - birth.getFullYear()
  const monthDiff = today.getMonth() - birth.getMonth()
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birth.getDate())) {
    age--
  }
  return age
}

// Calculate average spend per visit
const averageSpend = computed(() => {
  if (!customer.value || customer.value.total_visit === 0) return 0
  return customer.value.total_spend / customer.value.total_visit
})

// Calculate loyalty tier
const loyaltyTier = computed(() => {
  if (!customer.value) return 'Standard'
  const spend = customer.value.total_spend
  if (spend >= 5000) return 'Platinum'
  if (spend >= 3000) return 'Gold'
  if (spend >= 1000) return 'Silver'
  return 'Standard'
})

// Get tier color
const getTierColor = (tier: string) => {
  switch (tier) {
    case 'Platinum': return 'from-purple-500 to-indigo-500'
    case 'Gold': return 'from-amber-400 to-orange-500'
    case 'Silver': return 'from-slate-400 to-slate-500'
    default: return 'from-slate-300 to-slate-400'
  }
}

const handleDelete = async (id: number) => {
  try {
    const result = await DeleteCustomers(id)
    console.log(result)
    toast.success("Customer Deleted Successfully!");
    router.push('/customers')
  } catch (error) {
    const message = error.response?.data?.error || "Failed to delete customer";
    console.error('Error deleting customer:', message)
    toast.error(message);
  }
}
</script>

<template>
  <div class="space-y-6">

    <!-- Back Button & Actions -->
    <div class="flex items-center justify-between">
      <button @click="$router.back()" class="flex items-center gap-2 px-4 py-2 hover:bg-white/80 rounded-xl transition-all text-slate-600">
        <ArrowLeft class="w-4 h-4" />
        <span class="font-medium">Back to Customers</span>
      </button>

      <div class="flex items-center gap-3">
        <button @click="$router.push(`/customers/edit/${customer?.id}`)" class="px-5 py-2.5 bg-white hover:bg-gray-100 border border-slate-200 text-slate-700 rounded-xl text-sm font-medium transition-all flex items-center gap-2">
          <Edit2 class="w-4 h-4" />
          <span>Edit</span>
        </button>
        <button @click="customer?.id && handleDelete(customer.id)" class="px-5 py-2.5 bg-red-50 hover:bg-red-100 border border-red-200 text-red-600 rounded-xl text-sm font-medium transition-all flex items-center gap-2">
          <Trash2 class="w-4 h-4" />
          <span>Delete</span>
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="bg-white/80 backdrop-blur-xl rounded-2xl border border-white/60 shadow-lg p-20">
      <div class="text-center">
        <div class="inline-flex items-center justify-center w-16 h-16 rounded-full bg-gradient-to-r from-rose-400 via-purple-400 to-indigo-400 mb-4 animate-pulse">
          <User class="w-8 h-8 text-white" />
        </div>
        <p class="text-slate-500 font-medium">Loading customer details...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-if="error" class="bg-red-50 border-l-4 border-red-500 rounded-xl p-5 shadow-lg">
      <div class="flex items-start gap-3">
        <div class="w-10 h-10 rounded-xl bg-red-100 flex items-center justify-center flex-shrink-0">
          <span class="text-red-600 text-xl">⚠️</span>
        </div>
        <div>
          <p class="font-bold text-red-800 mb-1">Error Loading Customer</p>
          <p class="text-red-700 text-sm">{{ error }}</p>
        </div>
      </div>
    </div>

    <!-- Customer Details -->
    <div v-if="customer && !loading" class="space-y-6">

      <!-- Profile Header Card -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl border border-white/60 shadow-lg overflow-hidden">
        <!-- Cover Gradient -->
        <div class="h-32 bg-gradient-to-r from-rose-400 via-purple-400 to-indigo-400 relative">
          <div class="absolute inset-0 bg-black/10"></div>
        </div>

        <!-- Profile Info -->
        <div class="px-8 pb-8">
          <div class="flex flex-col lg:flex-row items-start lg:items-end gap-6 -mt-16">
            <!-- Avatar -->
            <div class="relative">
              <div class="w-32 h-32 rounded-3xl bg-gradient-to-br from-rose-400 via-purple-400 to-indigo-400 flex items-center justify-center text-white font-bold text-4xl shadow-2xl border-4 border-white">
                {{ getInitials(customer.name) }}
              </div>
              <!-- Loyalty Badge -->
              <div :class="['absolute -bottom-2 -right-2 px-4 py-1.5 rounded-xl shadow-lg text-white text-xs font-bold bg-gradient-to-r', getTierColor(loyaltyTier)]">
                {{ loyaltyTier }}
              </div>
            </div>

            <!-- Name & Info -->
            <div class="flex-1">
              <div class="flex items-start justify-between">
                <div>
                  <h1 class="text-3xl font-bold text-slate-800 mb-2">{{ customer.name }}</h1>
                  <p class="text-slate-500 text-sm mb-4">Customer ID: #{{ customer.id }}</p>

                  <div class="flex flex-wrap items-center gap-4 text-sm">
                    <div class="flex items-center gap-2 text-slate-600">
                      <Phone class="w-4 h-4 text-slate-400" />
                      <span>{{ customer.phone }}</span>
                    </div>
                    <div class="flex items-center gap-2 text-slate-600">
                      <Calendar class="w-4 h-4 text-slate-400" />
                      <span>{{ calculateAge(customer.birth_date) }} years old</span>
                    </div>
                    <div class="flex items-center gap-2 text-slate-600">
                      <User class="w-4 h-4 text-slate-400" />
                      <span class="capitalize">{{ customer.gender }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Stats Row -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
        <!-- Total Spend -->
        <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
          <div class="flex items-center gap-3 mb-2">
            <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-emerald-400 to-teal-500 flex items-center justify-center shadow-lg shadow-emerald-300/30">
              <DollarSign class="w-6 h-6 text-white" />
            </div>
            <div>
              <p class="text-xs text-slate-500 font-medium">Total Spend</p>
              <p class="text-2xl font-bold text-slate-800">{{ formatCurrency(customer.total_spend) }}</p>
            </div>
          </div>
        </div>

        <!-- Total Visits -->
        <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
          <div class="flex items-center gap-3 mb-2">
            <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-blue-400 to-indigo-500 flex items-center justify-center shadow-lg shadow-blue-300/30">
              <Activity class="w-6 h-6 text-white" />
            </div>
            <div>
              <p class="text-xs text-slate-500 font-medium">Total Visits</p>
              <p class="text-2xl font-bold text-slate-800">{{ customer.total_visit }}</p>
            </div>
          </div>
        </div>

        <!-- Average Spend -->
        <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
          <div class="flex items-center gap-3 mb-2">
            <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-purple-400 to-pink-500 flex items-center justify-center shadow-lg shadow-purple-300/30">
              <TrendingUp class="w-6 h-6 text-white" />
            </div>
            <div>
              <p class="text-xs text-slate-500 font-medium">Avg per Visit</p>
              <p class="text-2xl font-bold text-slate-800">{{ formatCurrency(averageSpend) }}</p>
            </div>
          </div>
        </div>

        <!-- Member Since -->
        <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-5 border border-white/60 shadow-lg">
          <div class="flex items-center gap-3 mb-2">
            <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-rose-400 to-orange-500 flex items-center justify-center shadow-lg shadow-rose-300/30">
              <Clock class="w-6 h-6 text-white" />
            </div>
            <div>
              <p class="text-xs text-slate-500 font-medium">Member Since</p>
              <p class="text-sm font-bold text-slate-800">{{ formatDate(customer.created_at).split(',')[0] }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Details Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">

        <!-- Personal Information -->
        <div class="lg:col-span-2 space-y-6">

          <!-- Contact & Personal Details -->
          <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
            <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
              <User class="w-5 h-5 text-purple-500" />
              Personal Information
            </h2>

            <div class="space-y-4">
              <div class="flex items-start gap-3 p-4 bg-slate-50 rounded-xl">
                <Phone class="w-5 h-5 text-slate-400 mt-0.5" />
                <div>
                  <p class="text-xs text-slate-500 font-medium mb-1">Phone Number</p>
                  <p class="text-sm font-semibold text-slate-800">{{ customer.phone }}</p>
                </div>
              </div>

              <div class="flex items-start gap-3 p-4 bg-slate-50 rounded-xl">
                <Calendar class="w-5 h-5 text-slate-400 mt-0.5" />
                <div>
                  <p class="text-xs text-slate-500 font-medium mb-1">Date of Birth</p>
                  <p class="text-sm font-semibold text-slate-800">{{ formatDate(customer.birth_date) }} ({{ calculateAge(customer.birth_date) }} years old)</p>
                </div>
              </div>

              <div class="flex items-start gap-3 p-4 bg-slate-50 rounded-xl">
                <User class="w-5 h-5 text-slate-400 mt-0.5" />
                <div>
                  <p class="text-xs text-slate-500 font-medium mb-1">Gender</p>
                  <p class="text-sm font-semibold text-slate-800 capitalize">{{ customer.gender }}</p>
                </div>
              </div>

              <div class="flex items-start gap-3 p-4 bg-slate-50 rounded-xl">
                <MapPin class="w-5 h-5 text-slate-400 mt-0.5" />
                <div>
                  <p class="text-xs text-slate-500 font-medium mb-1">Address</p>
                  <p class="text-sm font-semibold text-slate-800">{{ customer.address || 'No address provided' }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Notes -->
          <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
            <h2 class="text-lg font-semibold text-slate-800 mb-4 flex items-center gap-2">
              <FileText class="w-5 h-5 text-purple-500" />
              Additional Notes
            </h2>

            <div v-if="customer.note" class="p-4 bg-amber-50 border border-amber-200 rounded-xl">
              <p class="text-sm text-slate-700 whitespace-pre-wrap">{{ customer.note }}</p>
            </div>
            <div v-else class="p-8 text-center">
              <div class="inline-flex items-center justify-center w-12 h-12 rounded-full bg-slate-100 mb-3">
                <FileText class="w-6 h-6 text-slate-400" />
              </div>
              <p class="text-sm text-slate-500">No notes available for this customer</p>
            </div>
          </div>

        </div>

        <!-- Sidebar -->
        <div class="space-y-6">

          <!-- Loyalty Info -->
          <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
            <h2 class="text-lg font-semibold text-slate-800 mb-4 flex items-center gap-2">
              <Star class="w-5 h-5 text-purple-500" />
              Loyalty Status
            </h2>

            <div :class="['p-4 rounded-xl text-center bg-gradient-to-br', getTierColor(loyaltyTier)]">
              <div class="w-16 h-16 mx-auto mb-3 rounded-full bg-white/20 flex items-center justify-center">
                <Star class="w-8 h-8 text-white" />
              </div>
              <h3 class="text-2xl font-bold text-white mb-1">{{ loyaltyTier }}</h3>
              <p class="text-sm text-white/80">Member Tier</p>
            </div>

            <div class="mt-4 space-y-2 text-sm">
              <div class="flex justify-between items-center p-3 bg-slate-50 rounded-lg">
                <span class="text-slate-600">Lifetime Value</span>
                <span class="font-bold text-slate-800">{{ formatCurrency(customer.total_spend) }}</span>
              </div>
              <div class="flex justify-between items-center p-3 bg-slate-50 rounded-lg">
                <span class="text-slate-600">Next Tier</span>
                <span class="font-bold text-slate-800">
                  {{ loyaltyTier === 'Platinum' ? 'Max Tier' : loyaltyTier === 'Gold' ? '$5,000' : loyaltyTier === 'Silver' ? '$3,000' : '$1,000' }}
                </span>
              </div>
            </div>
          </div>

          <!-- Quick Actions -->
          <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
            <h2 class="text-lg font-semibold text-slate-800 mb-4 flex items-center gap-2">
              <Gift class="w-5 h-5 text-purple-500" />
              Quick Actions
            </h2>

            <div class="space-y-3">
              <button class="w-full px-4 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl transition-all">
                Book Appointment
              </button>
              <button class="w-full px-4 py-3 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-xl text-sm font-medium transition-all">
                View History
              </button>
              <button class="w-full px-4 py-3 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-xl text-sm font-medium transition-all">
                Send Message
              </button>
            </div>
          </div>

          <!-- Member Since -->
          <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
            <h2 class="text-lg font-semibold text-slate-800 mb-4 flex items-center gap-2">
              <Clock class="w-5 h-5 text-purple-500" />
              Membership
            </h2>

            <div class="text-center p-4 bg-gradient-to-br from-slate-50 to-blue-50 rounded-xl">
              <p class="text-sm text-slate-500 mb-1">Member Since</p>
              <p class="text-xl font-bold text-slate-800">{{ formatDate(customer.created_at) }}</p>
            </div>
          </div>

        </div>
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
