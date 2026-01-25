// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react'

const CommentIndicatorIcon = (): JSX.Element => (
    <svg
        className='CommentIndicatorIcon'
        width='16'
        height='16'
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
        aria-hidden='true'
        focusable='false'
    >
        <rect
            x='2.5'
            y='4'
            width='11'
            height='1.5'
            rx='0.75'
            fill='currentColor'
        />
        <rect
            x='2.5'
            y='7.25'
            width='11'
            height='1.5'
            rx='0.75'
            fill='currentColor'
        />
        <rect
            x='2.5'
            y='10.5'
            width='7.5'
            height='1.5'
            rx='0.75'
            fill='currentColor'
        />
    </svg>
)

export default CommentIndicatorIcon
