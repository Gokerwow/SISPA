// src/utils/date.ts
export const formatDate = (dateString: string | null) => {
  if (!dateString) return '-'

  const date = new Date(dateString)
  return new Intl.DateTimeFormat('en-GB', {
    day: 'numeric',
    month: 'short',
    year: 'numeric'
  }).format(date)
}

export const formatForInput = (dateString: string | null) => {
  if (!dateString) return ''

  // This takes "2026-02-13T01:39:16Z" and cuts it to "2026-02-13"
  return new Date(dateString).toISOString().split('T')[0]
}

