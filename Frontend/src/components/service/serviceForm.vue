<script setup lang="ts">
import { reactive, watch } from 'vue'
import {
  Heart,
  DollarSign,
  Save,
  X,
  Sparkles,
  Users,
  Globe
} from 'lucide-vue-next'
import type { Services } from '@/types/services'

// Form data
const formData = reactive<Services>({
  id: 0,
  name: '',
  price_male: 0,
  price_female: 0,
  price_tourist: 0,
  is_active: true,
  created_at: '',
})

const props = defineProps<{
  initialData?: Services,
  isSubmitting?: boolean
}>()

const emits = defineEmits<{
  (e: 'submit', data: Services): void;
}>()

watch(() => props.initialData, (newData) => {
  if (newData) {
    Object.assign(formData, newData)
  }
}, { immediate: true })

// Submit form
const handleSubmit = async () => {
  emits('submit', { ...formData })
}


</script>

<template>
  <div class="space-y-6">

    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1
          class="text-3xl font-bold bg-gradient-to-r from-rose-600 via-purple-600 to-indigo-600 bg-clip-text text-transparent mb-2">
          Add New Service
        </h1>
        <p class="text-slate-500 text-sm">Create a new spa service offering</p>
      </div>
      <button class="p-3 hover:bg-slate-100 rounded-xl transition-all text-slate-600">
        <X class="w-5 h-5" />
      </button>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">

      <!-- Service Information -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <Heart class="w-5 h-5 text-purple-500" />
          Service Information
        </h2>

        <div class="space-y-5">
          <!-- Service Name -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Service Name <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <Sparkles class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
              <input v-model="formData.name" type="text" required
                placeholder="e.g., Swedish Massage, Facial Treatment, Manicure"
                class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all" />
            </div>
            <p class="text-xs text-slate-500 mt-1.5">Enter a clear, descriptive name for your service</p>
          </div>
        </div>
      </div>

      <!-- Pricing Structure -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <DollarSign class="w-5 h-5 text-purple-500" />
          Pricing Structure
        </h2>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-5">

          <!-- Male Price -->
          <div class="bg-gradient-to-br from-blue-50 to-indigo-50 border-2 border-blue-200 rounded-xl p-5">
            <div class="flex items-center gap-2 mb-3">
              <div
                class="w-8 h-8 rounded-lg bg-gradient-to-br from-blue-400 to-indigo-500 flex items-center justify-center shadow-md">
                <Users class="w-4 h-4 text-white" />
              </div>
              <label class="text-sm font-semibold text-slate-800">
                Male Price <span class="text-red-500">*</span>
              </label>
            </div>
            <div class="relative">
              <span class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-500 font-medium">Rp</span>
              <input v-model.number="formData.price_male" type="number" required min="0" step="1000" placeholder="0"
                class="w-full pl-10 pr-4 py-3 bg-white border-2 border-blue-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-400 transition-all font-semibold text-slate-800" />
            </div>
            <p class="text-xs text-slate-600 mt-2">Price for male customers</p>
          </div>

          <!-- Female Price -->
          <div class="bg-gradient-to-br from-pink-50 to-rose-50 border-2 border-pink-200 rounded-xl p-5">
            <div class="flex items-center gap-2 mb-3">
              <div
                class="w-8 h-8 rounded-lg bg-gradient-to-br from-pink-400 to-rose-500 flex items-center justify-center shadow-md">
                <Users class="w-4 h-4 text-white" />
              </div>
              <label class="text-sm font-semibold text-slate-800">
                Female Price <span class="text-red-500">*</span>
              </label>
            </div>
            <div class="relative">
              <span class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-500 font-medium">Rp</span>
              <input v-model.number="formData.price_female" type="number" required min="0" step="1000" placeholder="0"
                class="w-full pl-10 pr-4 py-3 bg-white border-2 border-pink-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-pink-400 transition-all font-semibold text-slate-800" />
            </div>
            <p class="text-xs text-slate-600 mt-2">Price for female customers</p>
          </div>

          <!-- Tourist Price -->
          <div class="bg-gradient-to-br from-purple-50 to-indigo-50 border-2 border-purple-200 rounded-xl p-5">
            <div class="flex items-center gap-2 mb-3">
              <div
                class="w-8 h-8 rounded-lg bg-gradient-to-br from-purple-400 to-indigo-500 flex items-center justify-center shadow-md">
                <Globe class="w-4 h-4 text-white" />
              </div>
              <label class="text-sm font-semibold text-slate-800">
                Tourist Price <span class="text-red-500">*</span>
              </label>
            </div>
            <div class="relative">
              <span class="absolute left-4 top-1/2 -translate-y-1/2 text-slate-500 font-medium">Rp</span>
              <input v-model.number="formData.price_tourist" type="number" required min="0" step="1000" placeholder="0"
                class="w-full pl-10 pr-4 py-3 bg-white border-2 border-purple-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 transition-all font-semibold text-slate-800" />
            </div>
            <p class="text-xs text-slate-600 mt-2">Price for tourist customers</p>
          </div>
        </div>

        <!-- Pricing Tips -->
        <div class="mt-5 p-4 bg-amber-50 border border-amber-200 rounded-xl">
          <div class="flex items-start gap-3">
            <span class="text-xl">💡</span>
            <div>
              <p class="text-sm font-semibold text-amber-800 mb-1">Pricing Tips</p>
              <p class="text-xs text-amber-700">Consider your costs, competitor pricing, and target market. Tourist
                prices are typically higher to account for premium service expectations.</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Service Status -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <Sparkles class="w-5 h-5 text-purple-500" />
          Service Status
        </h2>

        <div class="space-y-3">
          <label class="flex items-center gap-4 p-5 rounded-xl border-2 cursor-pointer transition-all"
            :class="formData.is_active ? 'border-emerald-300 bg-emerald-50' : 'border-slate-200 bg-slate-50 hover:border-slate-300'">
            <input v-model="formData.is_active" type="radio" :value="true"
              class="w-5 h-5 text-emerald-500 focus:ring-emerald-400" />
            <div class="flex-1">
              <p class="font-semibold text-slate-800 flex items-center gap-2">
                Active
                <span class="px-2 py-0.5 bg-emerald-200 text-emerald-700 rounded-full text-xs">Available Now</span>
              </p>
              <p class="text-xs text-slate-600 mt-1">Service is available for booking</p>
            </div>
          </label>

          <label class="flex items-center gap-4 p-5 rounded-xl border-2 cursor-pointer transition-all"
            :class="!formData.is_active ? 'border-slate-300 bg-slate-50' : 'border-slate-200 bg-white hover:border-slate-300'">
            <input v-model="formData.is_active" type="radio" :value="false"
              class="w-5 h-5 text-slate-500 focus:ring-slate-400" />
            <div class="flex-1">
              <p class="font-semibold text-slate-800 flex items-center gap-2">
                Inactive
                <span class="px-2 py-0.5 bg-slate-200 text-slate-600 rounded-full text-xs">Hidden</span>
              </p>
              <p class="text-xs text-slate-600 mt-1">Service will not be available for booking</p>
            </div>
          </label>
        </div>
      </div>

      <!-- Price Preview -->
      <div class="bg-gradient-to-br from-purple-50 to-indigo-50 border-2 border-purple-200 rounded-2xl p-6 shadow-lg">
        <h3 class="text-sm font-semibold text-slate-800 mb-4 flex items-center gap-2">
          <DollarSign class="w-4 h-4 text-purple-600" />
          Price Summary
        </h3>
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
          <div class="bg-white rounded-xl p-4 border border-blue-200">
            <p class="text-xs text-slate-500 mb-1">Male</p>
            <p class="text-2xl font-bold text-slate-800">
              {{ formData.price_male > 0 ? `Rp ${formData.price_male.toLocaleString('id-ID')}` : 'Rp 0' }}
            </p>
          </div>
          <div class="bg-white rounded-xl p-4 border border-pink-200">
            <p class="text-xs text-slate-500 mb-1">Female</p>
            <p class="text-2xl font-bold text-slate-800">
              {{ formData.price_female > 0 ? `Rp ${formData.price_female.toLocaleString('id-ID')}` : 'Rp 0' }}
            </p>
          </div>
          <div class="bg-white rounded-xl p-4 border border-purple-200">
            <p class="text-xs text-slate-500 mb-1">Tourist</p>
            <p class="text-2xl font-bold text-slate-800">
              {{ formData.price_tourist > 0 ? `Rp ${formData.price_tourist.toLocaleString('id-ID')}` : 'Rp 0' }}
            </p>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex items-center justify-end gap-4">
        <button type="button"
          class="px-6 py-3 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-2xl text-sm font-medium transition-all flex items-center gap-2">
          <X class="w-4 h-4" />
          Cancel
        </button>

        <button type="submit" :disabled="isSubmitting"
          class="px-8 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl hover:shadow-purple-300/50 transition-all flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
          <Save class="w-4 h-4" />
          <span v-if="!isSubmitting">Save Service</span>
          <span v-else>Saving...</span>
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

/* Custom radio button */
input[type="radio"]:checked {
  background-color: rgb(16, 185, 129);
  border-color: rgb(16, 185, 129);
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
