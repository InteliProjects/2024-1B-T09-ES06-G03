/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './App.{js,jsx,ts,tsx}',
    './components/**/*.{js,jsx,ts,tsx}',
    './screens/**/*.{js,jsx,ts,tsx}',
  ],
  theme: {
    extend: {
      dropShadow: {
        'shadow-gray': '0px 4px 8px 0px #B3B5B7;',
        'shadow-green-button': [
          '11px 16px 5px 0px rgba(0, 0, 0, 0.00)', 
          '7px 10px 5px 0px rgba(0, 0, 0, 0.03)', 
          '4px 6px 4px 0px rgba(0, 0, 0, 0.10)', 
          '2px 3px 3px 0px rgba(0, 0, 0, 0.17)', 
          '0px 1px 2px 0px rgba(0, 0, 0, 0.20)'
        ],
        'shadow-card': '0px 4px 8px 0px rgba(0, 0, 0, 0.25)',
      },
      colors: {
        'green-10': '#3A8A88',
        'green-2': '#1F5E5C',
        'gray-1': '#F6F6F6',
        'gray-2': '#E5E5E5',
        'gray-3': '#ABABAB',
        'yellow-1': '#E3B146',
        'red-1': '#BB3756',
        'gray-4': 'rgba(105, 105, 105, 1)'
        
      },
    },
  },
  plugins: [],
};