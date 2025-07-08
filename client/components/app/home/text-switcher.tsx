'use client';

import { motion, AnimatePresence } from 'framer-motion';
import { useEffect, useState } from 'react';

const texts = ['yapacagin-isi', 'boyle-linki', 'bu-problemi'];

export default function AnimatedTextSwitcher() {
  const [index, setIndex] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setIndex(i => (i + 1) % texts.length);
    }, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="grow overflow-hidden relative text-4xl max-md:text-2xl font-black whitespace-nowrap">
      <AnimatePresence mode="wait">
        <motion.div
          key={index}
          initial={{ y: -40, opacity: 0 }}
          animate={{ y: 0, opacity: 1 }}
          exit={{ y: 40, opacity: 0 }}
          transition={{ duration: 0.5 }}
          className="absolute bottom-0"
        >
          {texts[index]}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}
