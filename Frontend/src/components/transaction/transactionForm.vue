<script setup lang="ts">
import type { FullCustomer } from '@/types/customer';
import type { Services } from '@/types/services';
import type { Therapist } from '@/types/therapist';
import type { InvoiceData, TransactionRequest } from '@/types/transaction';
import { computed, reactive, ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import {
  FileText,
  Save,
  X,
  Calendar,
  DollarSign,
  Users,
  UserCircle,
  Heart,
  Plus,
  Trash2,
  ShoppingCart,
} from 'lucide-vue-next'

const router = useRouter()
const discountType = ref<'amount' | 'percentage'>('amount')
const discountInput = ref(0)

const formData = reactive<TransactionRequest>({
  nota_number: '',
  date: new Date().toISOString().split('T')[0] ?? '',
  status: 'pending',
  discount_amount: 0,
  sub_total: 0,
  grand_total: 0,
  items: [
    { customer_id: 0, therapist_id: 0, service_id: 0 }
  ]
})
const customers = ref<FullCustomer[]>([])
const services = ref<Services[]>([])
const therapists = ref<Therapist[]>([])

const props = defineProps<{
  initialDataInvoice?: InvoiceData,
  initialDataCustomer: FullCustomer[],
  initialDataService: Services[],
  initialDataTherapist: Therapist[],
  isSubmitting: boolean
}>()

const emits = defineEmits<{
  (e: 'submit', data: TransactionRequest): void
}>()

// Add new item
const addItem = () => {
  formData.items.push({
    customer_id: 0,
    therapist_id: 0,
    service_id: 0
  })
}

// Remove item
const removeItem = (index: number) => {
  if (formData.items.length > 1) {
    formData.items.splice(index, 1)
    calculateTotals()
  }
}

// Get service price (using male price as default)
const getServicePrice = (serviceId: number) => {
  const service = props.initialDataService?.find(s => s.id === serviceId)
  return service ? service.price_male : 0
}

// Calculate totals
const calculateTotals = () => {
  // Calculate sub total from all items
  formData.sub_total = formData.items.reduce((sum, item) => {
    return sum + getServicePrice(item.service_id)
  }, 0)

  // Calculate grand total (sub_total - discount)
  formData.grand_total = formData.sub_total - formData.discount_amount
}

// Calculate discount based on type and then totals
const calculateDiscountAndTotals = () => {
  // Calculate sub total first
  formData.sub_total = formData.items.reduce((sum, item) => {
    return sum + getServicePrice(item.service_id)
  }, 0)

  // Calculate discount based on type
  if (discountType.value === 'percentage') {
    formData.discount_amount = Math.round((formData.sub_total * discountInput.value) / 100)
  } else {
    formData.discount_amount = discountInput.value
  }

  // Calculate grand total
  formData.grand_total = formData.sub_total - formData.discount_amount
}

// Watch for changes in items or discount
const watchCalculation = computed(() => {
  calculateTotals()
  return true
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

watch(() => props.initialDataInvoice, (newData) => {
  if (newData) {
    Object.assign(formData, newData)
    discountInput.value = newData.discount_amount
  }
})

watch(() => props.initialDataCustomer, (newData) => {
  if (newData) {
    customers.value = newData
  }
})

watch(() => props.initialDataService, (newData) => {
  if (newData) {
    services.value = newData
  }
})

watch(() => props.initialDataTherapist, (newData) => {
  if (newData) {
    therapists.value = newData
  }
})

// Submit form
const handleSubmit = () => {
  emits('submit', {...formData})
}

// Status options
const statusOptions = [
  { value: 'pending', label: 'Pending', color: 'amber' },
  { value: 'paid', label: 'Paid', color: 'emerald' },
  { value: 'cancelled', label: 'Cancelled', color: 'red' }
]

</script>

<template>
  <div class="space-y-6">

    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1
          class="text-3xl font-bold bg-gradient-to-r from-rose-600 via-purple-600 to-indigo-600 bg-clip-text text-transparent mb-2">
          Create New Transaction
        </h1>
        <p class="text-slate-500 text-sm">Record a new transaction with multiple services</p>
      </div>
      <button @click="router.back()" class="p-3 hover:bg-slate-100 rounded-xl transition-all text-slate-600">
        <X class="w-5 h-5" />
      </button>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">

      <!-- Transaction Information -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <FileText class="w-5 h-5 text-purple-500" />
          Transaction Information
        </h2>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          <!-- Invoice Number -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Invoice Number <span class="text-slate-400">(Optional)</span>
            </label>
            <input v-model="formData.nota_number" type="text" placeholder="Auto-generated if left empty"
              class="w-full px-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all" />
          </div>

          <!-- Date -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Transaction Date <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <Calendar class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
              <input v-model="formData.date" type="date" required
                class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all" />
            </div>
          </div>

          <!-- Status -->
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-slate-700 mb-3">
              Payment Status <span class="text-red-500">*</span>
            </label>
            <div class="flex gap-3">
              <label v-for="option in statusOptions" :key="option.value" :class="[
                'flex-1 flex items-center justify-center gap-2 p-3 rounded-xl border-2 cursor-pointer transition-all',
                formData.status === option.value
                  ? `border-${option.color}-300 bg-${option.color}-50`
                  : 'border-slate-200 bg-white hover:border-slate-300'
              ]">
                <input v-model="formData.status" type="radio" :value="option.value" class="w-4 h-4 text-purple-500" />
                <span class="font-medium text-sm capitalize">{{ option.label }}</span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- Transaction Items -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <div class="flex items-center justify-between mb-5">
          <h2 class="text-lg font-semibold text-slate-800 flex items-center gap-2">
            <ShoppingCart class="w-5 h-5 text-purple-500" />
            Transaction Items
          </h2>
          <button type="button" @click="addItem"
            class="px-4 py-2 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl transition-all flex items-center gap-2">
            <Plus class="w-4 h-4" />
            Add Item
          </button>
        </div>

        <div class="space-y-4">
          <div v-for="(item, index) in formData.items" :key="index"
            class="p-5 bg-gradient-to-br from-slate-50 to-purple-50 border-2 border-slate-200 rounded-xl">
            <div class="flex items-start justify-between mb-4">
              <h3 class="font-semibold text-slate-800">Item {{ index + 1 }}</h3>
              <button v-if="formData.items.length > 1" type="button" @click="removeItem(index)"
                class="p-2 hover:bg-red-100 text-red-600 rounded-lg transition-colors" title="Remove item">
                <Trash2 class="w-4 h-4" />
              </button>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <!-- Customer -->
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-2">
                  Customer <span class="text-red-500">*</span>
                </label>
                <div class="relative">
                  <Users class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 z-10" />
                  <select v-model="item.customer_id" @change="calculateTotals" required
                    class="w-full pl-10 pr-4 py-3 bg-white border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 transition-all appearance-none">
                    <option :value="0" disabled>Select customer</option>
                    <option v-for="customer in customers" :key="customer.id" :value="customer.id">
                      {{ customer.name }}
                    </option>
                  </select>
                </div>
              </div>

              <!-- Therapist -->
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-2">
                  Therapist <span class="text-red-500">*</span>
                </label>
                <div class="relative">
                  <UserCircle class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 z-10" />
                  <select v-model="item.therapist_id" @change="calculateTotals" required
                    class="w-full pl-10 pr-4 py-3 bg-white border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 transition-all appearance-none">
                    <option :value="0" disabled>Select therapist</option>
                    <option v-for="therapist in therapists" :key="therapist.id" :value="therapist.id">
                      {{ therapist.name }}
                    </option>
                  </select>
                </div>
              </div>

              <!-- Service -->
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-2">
                  Service <span class="text-red-500">*</span>
                </label>
                <div class="relative">
                  <Heart class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 z-10" />
                  <select v-model="item.service_id" @change="calculateTotals" required
                    class="w-full pl-10 pr-4 py-3 bg-white border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 transition-all appearance-none">
                    <option :value="0" disabled>Select service</option>
                    <option v-for="service in services" :key="service.id" :value="service.id">
                      {{ service.name }} - {{ formatCurrency(service.price_male) }}
                    </option>
                  </select>
                </div>
              </div>
            </div>

            <!-- Item Price Display -->
            <div v-if="item.service_id" class="mt-3 p-3 bg-white rounded-lg border border-purple-200">
              <div class="flex justify-between items-center">
                <span class="text-sm text-slate-600">Item Price:</span>
                <span class="text-lg font-bold text-purple-600">{{ formatCurrency(getServicePrice(item.service_id))
                  }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Pricing Calculation -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <DollarSign class="w-5 h-5 text-purple-500" />
          Pricing Summary
        </h2>

        <div class="space-y-4">
          <!-- Discount -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-3">
              Discount
            </label>

            <!-- Discount Type Toggle -->
            <div class="flex gap-2 mb-3">
              <button type="button" @click="discountType = 'amount'; calculateDiscountAndTotals()" :class="[
                'flex-1 py-2 px-4 rounded-xl text-sm font-medium transition-all',
                discountType === 'amount'
                  ? 'bg-gradient-to-r from-purple-500 to-indigo-500 text-white shadow-lg'
                  : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
              ]">
                Fixed Amount (Rp)
              </button>
              <button type="button" @click="discountType = 'percentage'; calculateDiscountAndTotals()" :class="[
                'flex-1 py-2 px-4 rounded-xl text-sm font-medium transition-all',
                discountType === 'percentage'
                  ? 'bg-gradient-to-r from-purple-500 to-indigo-500 text-white shadow-lg'
                  : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
              ]">
                Percentage (%)
              </button>
            </div>

            <!-- Discount Input -->
            <div class="relative">
              <span class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-500 font-medium">
                {{ discountType === 'amount' ? 'Rp' : '%' }}
              </span>

              <input v-model.number="discountInput" @input="calculateDiscountAndTotals" type="number" min="0"
                :max="discountType === 'percentage' ? 100 : undefined" :step="discountType === 'amount' ? 1000 : 1"
                placeholder="0"
                class="w-full pl-12 pr-16 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 transition-all font-semibold text-slate-800" />
            </div>

            <!-- Discount Amount Display (when using percentage) -->
            <p v-if="discountType === 'percentage' && discountInput > 0" class="mt-2 text-sm text-slate-600">
              Discount amount: <span class="font-bold text-purple-600">{{ formatCurrency(formData.discount_amount) }}</span>
            </p>
          </div>

          <!-- Summary -->
          <div class="p-5 bg-gradient-to-br from-slate-50 to-purple-50 rounded-xl border-2 border-purple-200">
            <!-- Ensure calculation runs -->
            <span v-if="watchCalculation" class="hidden"></span>

            <div class="space-y-3">
              <div class="flex justify-between items-center">
                <span class="text-slate-600 font-medium">Sub Total:</span>
                <span class="text-slate-800 font-semibold text-lg">{{ formatCurrency(formData.sub_total) }}</span>
              </div>
              <div class="flex justify-between items-center">
                <span class="text-slate-600 font-medium">Discount:</span>
                <span class="text-rose-600 font-semibold text-lg">-{{ formatCurrency(formData.discount_amount) }}</span>
              </div>
              <div class="flex justify-between items-center pt-3 border-t-2 border-purple-300">
                <span class="text-slate-800 font-bold text-xl">Grand Total:</span>
                <span :class="['font-bold text-3xl', formData.grand_total > 0 ? 'text-emerald-600' : 'text-red-600']">{{ formatCurrency(formData.grand_total) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex items-center justify-end gap-4">
        <button type="button" @click="router.back()"
          class="px-6 py-3 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-2xl text-sm font-medium transition-all flex items-center gap-2">
          <X class="w-4 h-4" />
          Cancel
        </button>

        <button type="submit" :disabled="isSubmitting"
          class="px-8 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl hover:shadow-purple-300/50 transition-all flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
          <Save class="w-4 h-4" />
          <span v-if="!isSubmitting">Create Transaction</span>
          <span v-else>Creating...</span>
        </button>
      </div>

    </form>
  </div>
</template>

<style scoped>
/* Smooth transitions */
* {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Remove number input arrows */
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
}

/* Smooth font rendering */
body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
