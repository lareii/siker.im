'use client';

import Navbar from '@/components/app/navbar';
import Footer from '@/components/app/home/footer';
import MainContent from '@/components/app/home/main-content';

import { motion } from 'framer-motion';

export default function Home() {
  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 0.5 }}
    >
      <Navbar />
      <MainContent />
      <Footer />
    </motion.div>
  );
}
