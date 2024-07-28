import React from 'react';

// Mapping of tags to emojis
const tagEmojiMap = {
  coffee: '☕️',
  breakfast: '🍳',
  lunch: '🥪',
  dinner: '🍽️',
  brunch: '🥞',
};

const TagEmoji = ({ tag }) => {
  // Get the emoji for the given tag, or a default if the tag is not found
  const emoji = tagEmojiMap[tag] || '❓';

  return (
    <div>
      <span className="size-28 text-opacity-25">
        <h6>{emoji}</h6>
      </span>
    </div>
  );
};

export default TagEmoji;
